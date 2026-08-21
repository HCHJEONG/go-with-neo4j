# goGraph

goGraph is an SSH-based Go and Neo4j system for reconstructing factory
operation conversations into evidence-grounded incident graphs.

## Local Development

Local development runs Go and Neo4j directly in WSL.

Confirmed baseline:

```text
Go 1.27.0
Neo4j 2026.07.1
cypher-shell 2026.07.1
Docker Engine 29.5.2 for later deployment packaging
Docker Compose v5.1.4 for later AWS runtime
```

Start Neo4j in one terminal:

```bash
sudo neo4j console
```

Verify Neo4j from another terminal:

```bash
cypher-shell -a neo4j://localhost:7687 -u neo4j -p '<password>' 'RETURN 1;'
```

Prepare local configuration:

```bash
cp .env.example .env
```

Edit `.env` and set `NEO4J_PASSWORD` to the local Neo4j password. Do not commit
`.env`.

Run tests:

```bash
go test ./...
```

Run the app skeleton:

```bash
set -a
. ./.env
set +a
go run ./cmd/gograph
```

## Current Phase 1 Boundary

The Neo4j Go driver dependency is intentionally not added yet. Add it manually
before implementing the real readiness check.

Do not implement SSH, parser, synthetic data generation, graph schema, vector
index, Docker Compose deployment, or AWS deployment in Phase 1.

