package main

import (
	"flag"

	"github.com/gofiber/fiber/v2/log"
	"github.com/somuthink/pics_journal/core/internal/config"
	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/handlers"
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

	handlers.Initialize()
}
