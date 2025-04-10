import json
from peft import LoraConfig, get_peft_model
from transformers import TrainingArguments, Trainer, AutoTokenizer, AutoModelForCausalLM, BitsAndBytesConfig
import os
import torch
from huggingface_hub import login, snapshot_download
from dotenv import load_dotenv, find_dotenv
from datasets import Dataset


def teaching():
    data = json.load(open('./dataset.json', 'r', encoding='UTF-8'))
    dataset = Dataset.from_list(data)
    load_dotenv(find_dotenv())
    login(token=os.environ.get("huggedface_key"))
    snapshot_download(repo_id='mistralai/Mistral-7B-Instruct-v0.1', ignore_patterns=['*.md'], local_dir='./mistral_model-inst',
                      token=os.environ.get("huggedface_key"))

    bnb_config = BitsAndBytesConfig(
        load_in_4bit=True,
        bnb_4bit_quant_type="nf4",
        bnb_4bit_compute_dtype=torch.float16,
    )

    model = AutoModelForCausalLM.from_pretrained("./mistral_model-inst",
                                                 device_map='auto',
                                                 quantization_config=bnb_config,
                                                 torch_dtype='auto')

    tokenizer = AutoTokenizer.from_pretrained('./mistral_model-inst')
    tokenizer.pad_token = tokenizer.eos_token

    def tokenizing(single_data):
        local_data = f'instruction: {single_data['instruction']}\ninput: {single_data['input']}\noutput: {str(single_data['output'])}'
        tokenized = tokenizer(
            local_data,
            truncation=True,
            padding="max_length",
            max_length=256,
            return_tensors="pt")

        return {"input_ids": tokenized["input_ids"].squeeze(0),
                "attention_mask": tokenized["attention_mask"].squeeze(0)}

    dataset = dataset.map(tokenizing, batched=False, remove_columns=dataset.column_names)

    lora_config = LoraConfig(
        r=8,
        lora_alpha=16,
        target_modules=["q_proj", "v_proj"],
        lora_dropout=0.05,
        bias="none",
        task_type="CAUSAL_LM"
    )

    model = get_peft_model(model, lora_config)
    model.print_trainable_parameters()

    training_args = TrainingArguments(
        output_dir="../models",
        per_device_train_batch_size=1,
        gradient_accumulation_steps=4,
        num_train_epochs=3,
        save_steps=100,
        logging_steps=10,
        learning_rate=2e-5,
        fp16=True
    )

    trainer = Trainer(
        model=model,
        args=training_args,
        train_dataset=dataset,
        data_collator=lambda d: {
            "input_ids": torch.stack([torch.tensor(d[0]["input_ids"]) for dat in d]),
            "attention_mask": torch.stack([torch.tensor(d[0]["attention_mask"]) for dat in d]),
            "labels": torch.stack([torch.tensor(d[0]["input_ids"]) for dat in d])
        }
    )

    trainer.train()
    model.save_pretrained("./lora-mistral-adapter-inst", safe_serialization=True)


teaching()
