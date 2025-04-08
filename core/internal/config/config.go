package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	VERSION   string
	JWT_TOKEN []byte
	DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
	SESSIONS_HOST, SESSIONS_PORT, SESSIONS_PASSWORD,
	INFERENCE_HOST,
	STORAGE_PREFIX string
	INFERENCE_PORT int
}

var Cfg Config

func Initialize(envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		return err
	}

	// version, err := strconv.atoi(os.getenv("VERSION"))
	// if err != nil {
	// 	return err
	// }

	inference_port, err := strconv.Atoi(os.Getenv("INFERENCE_PORT"))
	if err != nil {
		return err
	}

	Cfg = Config{
		VERSION: os.Getenv("VERSION"),

		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),

		SESSIONS_HOST:     os.Getenv("SESSIONS_HOST"),
		SESSIONS_PORT:     os.Getenv("SESSIONS_PORT"),
		SESSIONS_PASSWORD: os.Getenv("SESSIONS_PASSWORD"),

		INFERENCE_HOST: os.Getenv("INFERENCE_HOST"),
		INFERENCE_PORT: inference_port,

		STORAGE_PREFIX: os.Getenv("STORAGE_PREFIX"),

		JWT_TOKEN: []byte(os.Getenv("JWT_TOKEN")),
	}
	return nil
}
