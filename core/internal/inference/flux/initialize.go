package flux

import (
	"github.com/richinsley/comfy2go/client"
	"github.com/richinsley/comfy2go/graphapi"
	"github.com/somuthink/pics_journal/core/internal/config"
)

var (
	comfy        *client.ComfyClient
	defaultGraph *graphapi.Graph
)

func Initialize(workflowDir string) error {
	comfy = client.NewComfyClient(config.Cfg.INFERENCE_HOST, config.Cfg.INFERENCE_PORT, nil)

	if err := comfy.Init(); err != nil {
		return err
	}

	var err error

	defaultGraph, _, err = comfy.NewGraphFromJsonFile(workflowDir + "main.json")
	if err != nil {
		return err
	}

	return nil
}
