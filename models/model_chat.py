from transformers import AutoTokenizer, AutoModelForCausalLM, BitsAndBytesConfig
import torch
from peft import PeftModel


def mistral_chat(request):
    bnb_config = BitsAndBytesConfig(
        load_in_4bit=True,
        bnb_4bit_quant_type="nf4",
        bnb_4bit_compute_dtype=torch.float16,
    )

    tokenizer = AutoTokenizer.from_pretrained('./mistral_model-inst')
    model = AutoModelForCausalLM.from_pretrained("./mistral_model-inst",
                                                 device_map='auto',
                                                 quantization_config=bnb_config,
                                                 torch_dtype='auto')

    model = PeftModel.from_pretrained(model, "./lora-mistral-adapter-inst")

    inputs = tokenizer(request, return_tensors="pt").to(model.device)
    outputs = model.generate(**inputs, max_new_tokens=500)

    return tokenizer.decode(outputs[0])