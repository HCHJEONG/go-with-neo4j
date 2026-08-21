package config

import "testing"

func TestLoadReadsRequiredEnvironment(t *testing.T) {
	t.Setenv("SSH_USERNAME", "operator")
	t.Setenv("SSH_PASSWORD", "secret-password")
	t.Setenv("NEO4J_URI", "neo4j://localhost:7687")
	t.Setenv("NEO4J_USERNAME", "neo4j")
	t.Setenv("NEO4J_PASSWORD", "secret-password")
	t.Setenv("EMBEDDING_DIMENSIONS", "768")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.AppEnv != "development" {
		t.Fatalf("AppEnv = %q, want development", cfg.AppEnv)
	}
	if cfg.EmbeddingDimensions != 768 {
		t.Fatalf("EmbeddingDimensions = %d, want 768", cfg.EmbeddingDimensions)
	}
}

func TestLoadRejectsMissingRequiredEnvironment(t *testing.T) {
	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want missing required environment error")
	}
}

func TestLoadRejectsInvalidEmbeddingDimensions(t *testing.T) {
	t.Setenv("SSH_USERNAME", "operator")
	t.Setenv("SSH_PASSWORD", "secret-password")
	t.Setenv("NEO4J_URI", "neo4j://localhost:7687")
	t.Setenv("NEO4J_USERNAME", "neo4j")
	t.Setenv("NEO4J_PASSWORD", "secret-password")
	t.Setenv("EMBEDDING_DIMENSIONS", "not-a-number")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want invalid embedding dimensions error")
	}
}
