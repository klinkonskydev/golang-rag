package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL          string
	APIKey           string
	Model            string
	SystemPromptFile string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file, make sure this file exists")
	}
}

func Load() Config {
	baseURL, ok := os.LookupEnv("OPENAI_BASE_URL")
	if !ok {
		log.Fatal("OPENAI_BASE_URL wasn't found in your .env file.")
	}

	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		log.Fatal("OPENAI_API_KEY wasn't found in your .env file.")
	}

	model, ok := os.LookupEnv("OPENAI_MODEL")
	if !ok {
		log.Fatal("OPENAI_MODEL wasn't found in your .env file.")
	}

	systemPromptFile := os.Getenv("SYSTEM_PROMPT_FILE")

	cfg := Config{
		BaseURL: baseURL,
		APIKey: apiKey,
		Model: model,
		SystemPromptFile: systemPromptFile,
	}

	if baseURL == "" {
		cfg.BaseURL = "https://api.openapi.com/v1"
	}

	if model == "" {
		cfg.Model = "gpt-4o-mini"
	}

	return cfg
}
