package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/HCHJEONG/go-with-neo4j/internal/config"
	"github.com/HCHJEONG/go-with-neo4j/internal/graph"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("configuration invalid", "error", err)
		os.Exit(1)
	}

	slog.Info(
		"gograph configuration loaded",
		"app_env", cfg.AppEnv,
		"log_level", cfg.LogLevel,
		"ssh_listen_addr", cfg.SSHListenAddr,
		"neo4j_uri", cfg.Neo4jURI,
		"embedding_provider", cfg.EmbeddingProvider,
		"embedding_dimensions", cfg.EmbeddingDimensions,
	)

	if err := graph.CheckNeo4j(context.Background(), graph.Neo4jConfig{
		URI:      cfg.Neo4jURI,
		Username: cfg.Neo4jUsername,
		Password: cfg.Neo4jPassword,
		Timeout:  5 * time.Second,
	}); err != nil {
		slog.Error("neo4j connectivity unhealthy", "error", err)
		os.Exit(1)
	}

	slog.Info("neo4j connectivity healthy")
}
