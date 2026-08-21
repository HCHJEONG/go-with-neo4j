package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv              string
	LogLevel            string
	SSHListenAddr       string
	SSHUsername         string
	SSHPassword         string
	Neo4jURI            string
	Neo4jUsername       string
	Neo4jPassword       string
	EmbeddingProvider   string
	EmbeddingDimensions int
}

func Load() (Config, error) {
	cfg := Config{
		AppEnv:              getEnv("APP_ENV", "development"),
		LogLevel:            getEnv("LOG_LEVEL", "debug"),
		SSHListenAddr:       getEnv("SSH_LISTEN_ADDR", ":2222"),
		SSHUsername:         os.Getenv("SSH_USERNAME"),
		SSHPassword:         os.Getenv("SSH_PASSWORD"),
		Neo4jURI:            os.Getenv("NEO4J_URI"),
		Neo4jUsername:       os.Getenv("NEO4J_USERNAME"),
		Neo4jPassword:       os.Getenv("NEO4J_PASSWORD"),
		EmbeddingProvider:   getEnv("EMBEDDING_PROVIDER", "none"),
		EmbeddingDimensions: 384,
	}

	if raw := os.Getenv("EMBEDDING_DIMENSIONS"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("invalid EMBEDDING_DIMENSIONS: %w", err)
		}
		cfg.EmbeddingDimensions = n
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) Validate() error {
	var errs []error

	if c.SSHUsername == "" {
		errs = append(errs, errors.New("SSH_USERNAME is required"))
	}
	if c.SSHPassword == "" {
		errs = append(errs, errors.New("SSH_PASSWORD is required"))
	}
	if c.Neo4jURI == "" {
		errs = append(errs, errors.New("NEO4J_URI is required"))
	}
	if c.Neo4jUsername == "" {
		errs = append(errs, errors.New("NEO4J_USERNAME is required"))
	}
	if c.Neo4jPassword == "" {
		errs = append(errs, errors.New("NEO4J_PASSWORD is required"))
	}
	if c.EmbeddingDimensions <= 0 {
		errs = append(errs, errors.New("EMBEDDING_DIMENSIONS must be positive"))
	}

	return errors.Join(errs...)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
