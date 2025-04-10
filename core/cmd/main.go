package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2/log"
	"github.com/somuthink/pics_journal/core/internal/config"
	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/handlers"
	"github.com/somuthink/pics_journal/core/internal/inference/llm"
	"github.com/somuthink/pics_journal/core/internal/queue"
)

func main() {
	envPath := flag.String("env", "../../.env", "provide .env filepath")
	flag.Parse()

	if err := config.Initialize(*envPath); err != nil {
		log.Fatal(err)
	}

	if err := db.Initialize(); err != nil {
		log.Fatal("failed to init DB with", "err", err)
	}

	queue.Initialize()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go llm.StartWorker(ctx)

	handlers.Initialize()
}
