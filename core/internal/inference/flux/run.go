package flux

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/charmbracelet/log"
	"github.com/richinsley/comfy2go/client"
	"github.com/somuthink/pics_journal/core/internal/models"
)

func runDefault(job models.FluxJob) (string, error) {
	Initialize("../../workflows/")
	inputFile, err := os.Open("../static/images/uploads/" + job.InputName)
	if err != nil {
		return "", err
	}

	defer inputFile.Close()
	inputImage, _, err := image.Decode(inputFile)

	loadImageNode := defaultGraph.GetFirstNodeWithTitle("Load Image")
	if loadImageNode == nil {
		return "", fmt.Errorf("no loadimage node")
	} else {
		prop := loadImageNode.GetPropertyWithName("choose file to upload")
		if prop == nil {
			return "", fmt.Errorf("no load prop for the node")
		} else {
			uploadprop, _ := prop.ToImageUploadProperty()

			_, err := comfy.UploadImage(inputImage, "img", true, client.InputImageType, "", uploadprop)
			if err != nil {
				return "", fmt.Errorf("failed to upload image to node")
			}
		}
	}

	simple_api := defaultGraph.GetSimpleAPI(nil)
	positive := simple_api.Properties["Positive"]
	width := simple_api.Properties["Width"]
	height := simple_api.Properties["Height"]
	seed := simple_api.Properties["Seed"]
	steps := simple_api.Properties["Steps"]

	positive.SetValue("" + job.Prompt)
	width.SetValue(1280)
	height.SetValue(720)
	steps.SetValue(30)
	seed.SetValue(rand.Intn(1024717549447978))

	item, err := comfy.QueuePrompt(defaultGraph)
	if err != nil {
		return "", fmt.Errorf("no loadimage node")
	}

	timeout := time.After(4 * time.Minute)

	for continueLoop := true; continueLoop; {
		select {
		case msg := <-item.Messages:
			switch msg.Type {
			case "started":
				qm := msg.ToPromptMessageStarted()
				log.Info("started execution prompt", "id", qm.PromptID, "jobID", job.JobID)
			case "stopped":
				qm := msg.ToPromptMessageStopped()
				if qm.Exception != nil {
					return "", fmt.Errorf("error in execution: %s", *qm.Exception)
				}
				continueLoop = false
			case "data":
				qm := msg.ToPromptMessageData()
				for k, v := range qm.Data {
					if k == "images" || k == "gifs" {
						for _, output := range v {
							img_data, err := comfy.GetImage(output)
							if err != nil {
								log.Error("Failed to get image", "err", err)
								return "", fmt.Errorf("failed to get image: %w", err)
							}
							f, err := os.Create("../static/images/outputs/" + output.Filename)
							if err != nil {
								return "", fmt.Errorf("failed to create output file: %w", err)
							}
							_, err = f.Write(*img_data)
							if err != nil {
								return "", fmt.Errorf("failed to write image: %w", err)
							}
							f.Close()
							return output.Filename, nil
						}
					}
				}
			}
		case <-timeout:
			return "", fmt.Errorf("timeout waiting for completion")
		}
	}
	return "", nil
}
