# goGraph Development Plan

## 0. Development Environment Setup Notes

This section must be updated as the local environment is prepared.

Current local strategy:

- The repository is cloned under the user's WSL home-side development area.
- Go SDK is installed directly in WSL Ubuntu.
- Neo4j Community Edition is installed directly in WSL Ubuntu.
- VS Code will connect to this repository through WSL.
- Go development extensions and tools will be installed for the WSL environment.
- Docker is not required for normal local development.
- Docker is used later for deployment packaging.
- AWS runtime uses Docker Compose for the Go app and Neo4j together.

Current confirmed environment:

```bash
Repository path: /home/hchjeong/IntelliJProjects/go-with-neo4j
git@github.com:HCHJEONG/go-with-neo4j.git
Ubuntu: 22.04.5 LTS on WSL2
Git: 2.34.1
Go: 1.27.0
Docker Engine: 29.5.2
Docker Compose: v5.1.4
Neo4j: 2026.07.1
cypher-shell: 2026.07.1
Java: OpenJDK 21, used for local Neo4j runtime only
VS Code Remote: WSL Ubuntu-22.04
gopls: v0.23.0
dlv: 1.27.1
goimports: installed
staticcheck: 2026.2 (0.8.0)
air: 1.67.4
```

Completed setup:

1. Repository cloned under the WSL filesystem.
2. Git remote confirmed.
3. Go SDK installed manually from the official Linux tarball into `/usr/local/go`.
4. Go module initialized at the repository root:

```bash
go mod init github.com/HCHJEONG/go-with-neo4j
```

5. Docker Engine and Docker Compose verified for future deployment packaging:

```bash
docker version
docker compose version
docker run --rm hello-world
```

6. Neo4j apt repository added manually.
7. Neo4j Community Edition and `cypher-shell` installed through apt.
8. Neo4j started locally with:

```bash
sudo neo4j console
```

9. Local Bolt connectivity verified with `cypher-shell`.
10. `RETURN 1;` query succeeded.
11. VS Code opened against WSL Ubuntu-22.04.
12. Go VS Code extension installed and `gopls` language server is running.
13. Go development tools installed and verified:

```bash
gopls version
dlv version
goimports -help
staticcheck -version
air -v
```

Manual setup record:

The following steps were performed manually by the developer during initial
environment setup. This record intentionally describes the successful path only.

1. Confirmed the repository location and remote:

```bash
cd ~/IntelliJProjects/go-with-neo4j
pwd
git remote -v
```

Confirmed:

```text
/home/hchjeong/IntelliJProjects/go-with-neo4j
origin git@github.com:HCHJEONG/go-with-neo4j.git
```

2. Installed Go SDK manually from the official tarball:

```bash
cd /tmp
wget https://go.dev/dl/go1.27.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.27.0.linux-amd64.tar.gz
```

3. Added Go binaries to the shell path for the current shell:

```bash
export PATH="/usr/local/go/bin:$HOME/go/bin:$PATH"
```

The permanent shell configuration still needs to be confirmed.

4. Initialized the Go module from the repository root:

```bash
cd ~/IntelliJProjects/go-with-neo4j
go mod init github.com/HCHJEONG/go-with-neo4j
cat go.mod
```

Confirmed:

```go
module github.com/HCHJEONG/go-with-neo4j

go 1.27.0
```

5. Verified Docker for later deployment packaging:

```bash
docker version
docker compose version
docker run --rm hello-world
```

Confirmed:

```text
Docker Engine 29.5.2
Docker Compose v5.1.4
hello-world succeeded
```

6. Added the Neo4j apt repository and installed Neo4j manually:

```bash
wget -O - https://debian.neo4j.com/neotechnology.gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/neo4j.gpg
echo "deb [signed-by=/usr/share/keyrings/neo4j.gpg] https://debian.neo4j.com stable latest" | sudo tee /etc/apt/sources.list.d/neo4j.list
sudo apt update
sudo apt install neo4j
```

The Neo4j package installed `neo4j`, `cypher-shell`, and the Java runtime needed
by local Neo4j.

7. Verified Neo4j command-line tools:

```bash
neo4j --version
cypher-shell --version
```

Confirmed:

```text
Neo4j 2026.07.1
Cypher-Shell 2026.07.1
```

8. Started local Neo4j manually:

```bash
sudo neo4j console
```

Confirmed from Neo4j logs:

```text
Bolt enabled on localhost:7687
HTTP enabled on localhost:7474
Started
```

9. Verified local Neo4j connectivity from a second terminal:

```bash
cypher-shell -a neo4j://localhost:7687 -u neo4j -p neo4j
```

The first successful login required changing the default password. The new
password must be kept out of Git and later placed only in local `.env`.

10. Verified Cypher execution:

```cypher
RETURN 1;
:exit
```

Confirmed:

```text
1 row
```

11. Opened the repository in VS Code through WSL and installed the Go extension.

Confirmed:

```text
VS Code Remote: WSL Ubuntu-22.04
Go extension active
GOPATH=/home/hchjeong/go
GOBIN=/home/hchjeong/go/bin
gopls language server running
```

12. Installed and verified Go development tools:

```bash
go install github.com/air-verse/air@latest
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

`gopls` and `dlv` were installed through the VS Code Go tooling flow.

Verification:

```bash
gopls version
dlv version
goimports -help
staticcheck -version
air -v
```

Confirmed:

```text
gopls v0.23.0
dlv 1.27.1
goimports installed
staticcheck 2026.2 (0.8.0)
air 1.67.4
```

14. Permanent Go PATH configuration verified from WSL:

```bash
grep -n '/usr/local/go/bin' ~/.bashrc
grep -n '$HOME/go/bin' ~/.bashrc
which go
which gopls
which dlv
which goimports
which staticcheck
which air
go version
```

Confirmed:

```text
~/.bashrc contains: export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH
/usr/local/go/bin/go
/home/hchjeong/go/bin/gopls
/home/hchjeong/go/bin/dlv
/home/hchjeong/go/bin/goimports
/home/hchjeong/go/bin/staticcheck
/home/hchjeong/go/bin/air
go version go1.27.0 linux/amd64
```

15. Neo4j Go driver dependency was added manually after Phase 1A:

```bash
go get github.com/neo4j/neo4j-go-driver/v5
```

Current repository state after checking `go.mod`, `go.sum`, and `go list -m all`:

```text
go.mod contains github.com/neo4j/neo4j-go-driver/v5 v5.28.4 // indirect.
go.sum contains checksums for github.com/neo4j/neo4j-go-driver/v5 v5.28.4.
go list -m all shows github.com/neo4j/neo4j-go-driver/v5 v5.28.4.
```

The dependency is currently marked indirect because the application code has not
imported it yet. Phase 1B should implement the readiness check and then run
`go mod tidy`, which should make the dependency direct if it is imported by
production code.

16. Phase 1B Neo4j readiness code was added manually in VS Code.

Files edited:

```text
internal/graph/neo4j.go
cmd/gograph/main.go
```

The graph package now creates a Neo4j driver with context, applies a timeout,
calls `VerifyConnectivity`, and returns wrapped errors. The application entry
point now loads config and calls the readiness check.

17. Phase 1B module cleanup, formatting, and tests were run manually:

```bash
go mod tidy
goimports -w cmd internal
go test ./...
```

Confirmed:

```text
github.com/neo4j/neo4j-go-driver/v5 v5.28.4 is now a direct dependency.
go test ./... passed.
```

18. Phase 1B application run was verified manually with local `.env` loaded:

```bash
set -a
. ./.env
set +a
go run ./cmd/gograph
```

Confirmed:

```text
INFO gograph configuration loaded
INFO neo4j connectivity healthy
```

Remaining setup:

1. Install or enable VS Code Docker and YAML extensions if not already enabled.
2. Confirm VS Code uses the WSL Go toolchain from its integrated terminal.
3. Open the repository from WSL when needed with:

```bash
code .
```

4. Create local `.env` from `.env.example`.
5. Implement local Neo4j readiness check from Go and finalize the Neo4j driver
   dependency in `go.mod`.
6. Later, add Dockerfile and AWS Docker Compose deployment files.

Reference setup order:

1. Confirm the repository path under WSL.
2. Confirm Git remote.
3. Install Go SDK in WSL.
4. Install Go development tools: `gopls`, `dlv`, `air`, `goimports`, `staticcheck`.
5. Add `$HOME/go/bin` to `PATH`.
6. Install or enable VS Code WSL, Go, Docker, and YAML extensions.
7. Open the repository from WSL with `code .`.
8. Install Neo4j Community Edition in WSL.
9. Confirm `neo4j` and `cypher-shell`.
10. Initialize the Go module.
11. Create the basic project structure.
12. Implement local Neo4j readiness check from Go.
13. Later, add Dockerfile and AWS Docker Compose deployment files.

Prerequisite groups:

- Go development: Git, Go SDK, VS Code WSL integration, Go VS Code extension,
  `gopls`, `dlv`, `air`, `goimports`, `staticcheck`.
- Local Neo4j runtime: Neo4j Community Edition, `cypher-shell`, and Java runtime
  only if required by the selected Neo4j installation method.

Java is not part of the Go development environment. It is checked only because
Neo4j runs on the JVM when Neo4j is installed directly in WSL.

Confirmed decisions:

- Local development: WSL-native Go and WSL-native Neo4j.
- Remote deployment: Docker image built locally from a clean clone, transferred
  as a tar file, then run on AWS with Docker Compose.
- Initial setup, local execution, deployment packaging, and remote deployment
  are manual-first. Do not add automation until the manual runbook has been
  executed, understood, and documented.
- SSH MVP authentication: password-based, loaded from environment variables.
- Search: Neo4j vector index is included in the planned MVP search path.

## 0.1 Go Learning Strategy

The developer will approach Go as a modern C for backend systems programming.

Guidelines:

- Prefer structs, functions, and explicit composition.
- Treat `main` as the composition root where dependencies are wired manually.
- Avoid Spring-style dependency injection containers and annotation-driven
  wiring.
- Use small interfaces only where they simplify a real boundary.
- Treat explicit `error` returns like checked operational outcomes.
- Use `context.Context` for cancellation, timeout, and request lifetime.
- Think about value vs pointer semantics deliberately.
- Prefer the standard library before adding dependencies.
- Keep the result easy to build as a single deployable binary.

## 1. Purpose

This document is the implementation plan for goGraph.

goGraph is an SSH-based operational intelligence system that ingests large factory operation conversations, reconstructs them into incident-level knowledge, stores the relationships in Neo4j, and helps operators recall similar past incidents with evidence.

Codex must follow this document when implementing the project.

Core rules:

1. Implement one phase at a time.
2. Do not pull later-phase features into earlier phases.
3. Keep every important result traceable to source messages.
4. Do not let an LLM generate or execute raw Cypher.
5. Validate all extracted incident data before persistence.
6. Do not commit secrets, real private chats, or operational data.
7. Update tests and documentation with implementation changes.
8. Do not create AWS resources, push to a remote, or install system packages without user approval.
9. Prefer documented manual commands before scripts. Add scripts only after the manual process is proven.

## 2. Project

Project name: `goGraph`

Repository: `github.com/HCHJEONG/go-with-neo4j`

Go module path:

```bash
github.com/HCHJEONG/go-with-neo4j
```

One-line description:

```text
An SSH-based Go and Neo4j system that converts fragmented factory conversations into evidence-grounded incident graphs and retrieves similar past operational situations.
```

The system must answer questions such as:

- Was there a similar past incident?
- Which asset, order, product, or line was affected?
- What action was taken?
- Did the action resolve the issue?
- How long did first response and recovery take?
- Which original messages support the answer?

## 3. Technical Stack

Application:

- Go
- Go modules
- Neo4j official Go driver
- Neo4j vector index for semantic search
- Charmbracelet Wish for SSH
- Line-oriented terminal shell first
- `log/slog` for structured logging
- Environment variables for configuration
- Standard `testing`
- Air, Delve, goimports, Staticcheck for development

Database:

- Neo4j Community Edition
- Bolt protocol
- Local development uses Neo4j installed directly in WSL
- Remote deployment uses Neo4j through Docker Compose
- Persistent data directory or volume
- Idempotent schema and index initialization
- Full-text index for keyword baseline
- Vector index for incident/message embeddings

Deployment target:

- AWS EC2
- Docker Engine
- Docker Compose
- Go application container
- Neo4j container
- EBS-backed Neo4j data volume

Do not introduce Kubernetes, ECS, a web frontend, or complex user management before MVP completion.

## 4. Local Development Assumptions

Development environment:

```text
Windows
└── WSL2 Ubuntu
    ├── Git
    ├── Go SDK
    ├── Neo4j Community Edition
    ├── Air
    ├── Delve
    └── VS Code Server
```

Docker is not required for normal local development. Docker Compose is used for
the remote AWS runtime strategy, where the Go app and Neo4j run together as
containers.

The source code should remain inside the WSL filesystem, not under `/mnt/c`.

Before changing the system, Codex must check current state:

```bash
uname -a
cat /etc/os-release
go version
git --version
```

For local Neo4j readiness, check separately:

```bash
java -version
neo4j --version
cypher-shell --version
```

Check whether Neo4j is installed as a service or as a manually started process:

```bash
systemctl status neo4j --no-pager
service neo4j status
ps aux | grep neo4j
```

Only report missing tools. Do not reinstall tools that already work.

## 5. Repository Structure

Initial structure:

```text
go-with-neo4j/
├── cmd/
│   └── gograph/
│       └── main.go
├── internal/
│   ├── config/
│   ├── chat/
│   ├── incident/
│   ├── graph/
│   ├── search/
│   ├── synth/
│   └── ssh/
├── migrations/
├── datasets/
│   ├── samples/
│   └── generated/
├── scripts/
├── tests/
├── .air.toml
├── .dockerignore
├── .env.example
├── .gitignore
├── compose.production.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
├── PLAN.md
└── README.md
```

Prefer one binary with subcommands instead of multiple early binaries:

```bash
gograph serve
gograph ingest
gograph generate
gograph schema
```

Avoid early over-abstraction:

- Do not create repository interfaces unless there are multiple implementations.
- Do not add a dependency injection framework.
- Do not build a plugin system.
- Do not create a general DSL.
- Do not build future-only package layers.

## 6. Configuration

All runtime configuration must come from environment variables.

`.env.example`:

```dotenv
APP_ENV=development
LOG_LEVEL=debug

SSH_LISTEN_ADDR=:2222
SSH_USERNAME=operator
SSH_PASSWORD=replace-me

NEO4J_URI=neo4j://localhost:7687
NEO4J_USERNAME=neo4j
NEO4J_PASSWORD=replace-me

EMBEDDING_PROVIDER=none
EMBEDDING_DIMENSIONS=384
```

Rules:

- `.env` must be ignored by Git.
- Passwords must never be hardcoded in Go source.
- Logs must not print secrets.
- If required variables are empty, the app must fail fast with a clear error.

## 7. Local Neo4j

Local development runs both Go and Neo4j directly in WSL.

Expected local runtime:

```bash
neo4j console
```

or, if Neo4j is installed as a service:

```bash
sudo systemctl start neo4j
```

Local checks:

```bash
cypher-shell -a neo4j://localhost:7687 -u neo4j -p '<password>' 'RETURN 1'
```

Local rules:

- Do not require Docker to run Phase 1 locally.
- Keep Neo4j Bolt bound to localhost for development.
- Store the local Neo4j password in `.env`, not in source code.
- Document the exact local Neo4j install/start method after it is confirmed.
- Do not commit local Neo4j data.

Remote Compose requirements are defined in the AWS deployment phase.

Neo4j ports must not be exposed publicly in production.

## 8. Graph Model

Initial nodes:

- `Message`
- `Actor`
- `Incident`
- `Asset`
- `Order`
- `Product`
- `Location`
- `Issue`
- `Action`

Initial relationships:

```text
(Actor)-[:SENT]->(Message)
(Message)-[:PART_OF]->(Incident)
(Message)-[:MENTIONS]->(Asset)
(Incident)-[:AFFECTS]->(Asset)
(Incident)-[:AFFECTS]->(Order)
(Incident)-[:OCCURRED_AT]->(Location)
(Incident)-[:HAS_ISSUE]->(Issue)
(Incident)-[:RESPONDED_WITH]->(Action)
(Action)-[:PERFORMED_BY]->(Actor)
(Incident)-[:THREATENS]->(Order)
(Order)-[:FOR_PRODUCT]->(Product)
```

`SIMILAR_TO` should not be persisted in the MVP unless the score, method, timestamp, and evidence are also stored. Similarity can be computed at query time first.

## 9. Schema And Indexes

Schema initialization must be idempotent.

Minimum constraints:

- `Message.id` unique
- `Message.contentHash` indexed or unique per source strategy
- `Actor.id` unique
- `Incident.id` unique
- `Asset.id` unique
- `Order.id` unique
- `Product.id` unique
- `Issue.id` unique
- `Action.id` unique

Search indexes:

- Full-text index over incident summaries and message text.
- Vector index over incident embeddings.
- Optionally vector index over message embeddings after the incident-level path works.

Vector index rules:

- Embedding dimensions must be configured and validated.
- Embedding provider must be replaceable without changing persistence logic.
- If `EMBEDDING_PROVIDER=none`, the app must still run with keyword and graph search.
- Tests must not require an external paid embedding API.

## 10. Data Pipeline

```text
Raw TXT
  -> Streaming Parser
  -> Normalized Message
  -> Episode Segmentation
  -> Entity and Relation Extraction
  -> Incident IR
  -> Validation
  -> Neo4j Persistence
  -> Keyword and Vector Candidate Retrieval
  -> Graph Expansion
  -> Deterministic Time Analysis
  -> Evidence-grounded Output
```

The parser must stream input and handle:

- Korean UTF-8 text
- date separators
- participant names
- message times
- multi-line messages
- empty lines
- malformed lines
- duplicate imports
- very long messages
- system messages

Internal message type:

```go
type Message struct {
    ID         string
    Source     string
    SourceLine int
    SentAt     time.Time
    Actor      string
    Text       string
    Hash       string
}
```

Incident IR:

```go
type IncidentIR struct {
    ExternalID string
    Type       string
    Summary    string
    Severity   string
    StartedAt  *time.Time
    ResolvedAt *time.Time
    AssetIDs   []string
    OrderIDs   []string
    Actions    []ActionIR
    Evidence   []EvidenceRef
    Confidence float64
}
```

Validation must check:

- allowed incident type
- required fields
- valid confidence range
- logical time ordering
- existing evidence messages
- no unevidenced asset/order references
- no direct Cypher generated from model output

## 11. Synthetic Factory Conversation Data

This area is critical. Do not create random chat text first and then guess the truth.

Use a ground-truth-first generator.

### 11.1 Ground Truth

Generate structured incident truth first:

```json
{
  "incidentId": "INC-0042",
  "line": "LINE-A",
  "asset": "CV-SENSOR-07",
  "issue": "intermittent-stop",
  "severity": "high",
  "startedAt": "2026-08-21T14:10:00+09:00",
  "dueAt": "2026-08-21T18:00:00+09:00",
  "actions": [
    {
      "at": "2026-08-21T14:23:00+09:00",
      "type": "inspect",
      "actor": "maintenance-01"
    },
    {
      "at": "2026-08-21T14:41:00+09:00",
      "type": "replace",
      "actor": "maintenance-02"
    }
  ],
  "resolvedAt": "2026-08-21T15:02:00+09:00",
  "affectedOrders": ["ORD-1017"],
  "expectedEvidenceMessageKeys": []
}
```

Truth JSON must be saved next to the generated TXT.

### 11.2 Domain Catalogs

The generator must use small explicit catalogs:

- production lines
- assets
- asset aliases
- products
- orders
- customers
- actors
- actor roles
- issue types
- action types
- resolution outcomes
- common Korean shop-floor expressions

Asset aliases are required because real conversations rarely use canonical IDs every time.

Example:

```text
CV-SENSOR-07
aliases: 컨베이어 센서 7번, CV 센서, 7번 센서, 컨베어 센서
```

### 11.3 Conversation Rendering

Generate messages from truth using templates and controlled variation.

Each incident should include at least:

- first symptom report
- asset or line mention
- production/order pressure mention when applicable
- investigation or assignment message
- one or more action messages
- outcome or unresolved status message

Messages should be interleaved with:

- unrelated routine work
- shift handover messages
- inventory chatter
- quality checks
- short acknowledgements
- typos and abbreviations
- multi-line messages
- incomplete messages

### 11.4 Similarity Test Design

The synthetic set must include controlled pairs:

- same asset, same issue, same successful action
- same asset, same issue, different failed action
- same issue on different line
- same symptom but different root cause
- same keywords but unrelated incident
- no shared keywords but similar operational pattern

These pairs are required so vector search and graph ranking can be evaluated meaningfully.

### 11.5 Determinism

The generator must support:

```bash
gograph generate --seed 42 --incidents 20 --messages 1000 --out ./datasets/generated
```

Same seed and parameters must produce byte-stable output.

MVP target:

- 20 incidents
- 8 to 12 actors
- 10 assets
- 10 orders
- about 1,000 messages

Expansion target:

- 100 to 200 incidents
- 10,000+ messages
- overlapping incidents
- confusing similar incidents

### 11.6 Evaluation Files

Generated output must include:

```text
factory-chat.txt
ground-truth.json
catalog.json
expected-recall-cases.json
```

`expected-recall-cases.json` should define queries and expected top-k incident IDs.

### 11.7 Real Data Rule

Do not use real company conversations or personal chats.

Public maintenance datasets may be used only as vocabulary and scenario inspiration if their licenses allow it. README must cite any public source used.

## 12. Search Strategy

Search must combine multiple signals.

Stage 1: candidate retrieval

- full-text keyword search
- Neo4j vector index over incident summaries/embeddings

Stage 2: graph expansion

- same asset
- same issue
- same line
- same product/order pressure
- same action
- similar outcome
- adjacent timeline

Stage 3: deterministic time analysis in Go

- first response time
- recovery time
- downtime duration
- due-date remaining time
- action intervals

Stage 4: evidence-grounded output

Every important search result must include:

- incident ID
- summary
- related asset/order
- score breakdown
- timeline
- action and result
- evidence message IDs
- clear distinction between source fact and derived inference

## 13. SSH Interface

The product interface is SSH only for MVP.

Use password authentication for the MVP.

Rules:

- Username and password come from environment variables.
- Password must not be hardcoded.
- Password must not be logged.
- Add basic rate limiting or failed-attempt delay before production deployment.
- The app SSH port is separate from OS management SSH.

Initial interface is a line-oriented shell:

```text
$ ssh -p 2222 operator@localhost

GOGRAPH OPERATIONS TERMINAL
Database: connected
Incidents: 20

gograph>
```

Initial commands:

```text
help
status
ingest
incidents
open <incident-id>
recall <natural-language-query>
graph <incident-id>
explain <incident-id>
quit
```

One-shot SSH commands and stdin ingest may be added after the interactive shell works. Do not implement them in Phase 1.

## 14. Implementation Phases

### Phase 0. Environment Verification

Tasks:

- Check WSL and OS details.
- Check Go.
- Check Git.
- Check Java runtime if Neo4j requires it.
- Check local Neo4j installation.
- Check `cypher-shell`.
- Check development tools if present: Air, gopls, Delve, goimports, Staticcheck.
- Check Docker and Docker Compose for later deployment packaging.
- Separate installed, missing, and broken tools.

Completion criteria:

```bash
go version
gopls version
dlv version
goimports -help
staticcheck -version
air -v
git --version
java -version
neo4j --version
cypher-shell --version
docker version
docker compose version
```

Status: completed manually. All listed tools have been verified.

Do not install system packages without user approval.

### Phase 1. Repository And Local Neo4j Readiness

Scope is intentionally small.

Phase 1A tasks before adding the Neo4j Go driver:

- Initialize Go module with `github.com/HCHJEONG/go-with-neo4j`.
- Create the initial repository structure.
- Add `.gitignore`.
- Add `.env.example`.
- Document the confirmed local Neo4j install/start method.
- Start local Neo4j directly in WSL.
- Implement configuration loading.
- Implement a minimal app entrypoint.
- Add basic tests for config validation.
- Write README local setup instructions.

Phase 1A status: implemented. The app currently validates configuration and
logs that the Neo4j readiness check is pending manual driver dependency
installation.

Manual step before Phase 1B:

```bash
go get github.com/neo4j/neo4j-go-driver/v5
go mod tidy
```

Phase 1B tasks after the manual driver dependency step:

- Connect to Neo4j from Go.
- Implement timeout-bound readiness check.
- Update tests for readiness behavior where practical.
- Update README with the final Phase 1 run command.

Do not implement:

- SSH server
- parser
- synthetic data generator
- graph schema beyond minimal connection check
- search
- vector index
- Docker Compose deployment
- AWS deployment

Completion criteria:

```bash
cp .env.example .env
neo4j console
go run ./cmd/gograph
go test ./...
```

Phase 1A completion: `go test ./...` passes and `go run ./cmd/gograph` loads
configuration with environment variables supplied.

Full Phase 1 completion after Phase 1B: the app must log that Neo4j connectivity
is healthy.

### Phase 2. Graph Schema And Persistence

Goal:

Create the first durable graph model in Neo4j and prove that the application can
initialize schema, save a minimal incident with evidence messages, and read it
back.

Scope rules:

- Keep Phase 2 focused on deterministic Go code and fixed Cypher templates.
- Do not implement parser, synthetic generator, SSH commands, search ranking, or
  LLM extraction in this phase.
- Do not let generated text or user input become raw Cypher.
- Use Cypher parameters for all values.
- Schema initialization must be safe to run repeatedly.
- Use local WSL Neo4j, not Docker Compose.

Phase 2A. Schema initialization:

- Add idempotent schema initialization.
- Create uniqueness constraints for initial node IDs.
- Create useful indexes for `Message.contentHash`, incident status/type, and
  message timestamps.
- Create full-text index for incident summaries and message text if supported by
  the installed Neo4j version.
- Create vector index for incident embeddings with configured dimensions.
- Keep migration/init code explicit and inspectable.

Initial constraints:

```text
Message.id
Actor.id
Incident.id
Asset.id
Order.id
Product.id
Location.id
Issue.id
Action.id
```

Initial indexes:

```text
Message.contentHash
Message.sentAt
Incident.type
Incident.status
Incident.startedAt
```

Phase 2B. Minimal domain types:

- Add Go structs for `Message`, `Incident`, and `EvidenceRef`.
- Keep fields aligned with the planned graph model.
- Do not introduce broad repository interfaces yet.
- Prefer a concrete Neo4j store type such as `graph.Store`.

Phase 2C. Persistence:

- Save a minimal incident.
- Save evidence messages.
- Link messages to incidents with `(:Message)-[:PART_OF]->(:Incident)`.
- Save actor/message relation if actor is provided.
- Read an incident by ID.
- Read evidence messages for an incident.
- Make repeated save behavior deterministic.

Phase 2D. Tests:

- Add integration tests that use the local Neo4j instance only when explicitly
  enabled.
- Skip integration tests by default if required env vars are missing.
- Use a test label or test ID prefix so manual cleanup is easy.
- Verify schema initialization can run twice.
- Verify a test incident can be saved and read.
- Verify evidence messages can be traversed from incident.

Suggested integration-test environment:

```bash
GOGRAPH_INTEGRATION_TESTS=1
NEO4J_URI=neo4j://localhost:7687
NEO4J_USERNAME=neo4j
NEO4J_PASSWORD=<local-password>
```

Suggested commands:

```bash
set -a
. ./.env
set +a
go test ./...
GOGRAPH_INTEGRATION_TESTS=1 go test ./internal/graph -run Integration
```

Completion criteria:

- Schema initialization can run repeatedly.
- A test incident can be saved and read.
- Evidence messages can be traversed from incident.
- `go test ./...` passes without requiring Neo4j.
- Integration tests pass when explicitly enabled against local Neo4j.
- README documents how to run Phase 2 schema and integration checks.

### Phase 3. Message Parser

Tasks:

- Implement Kakao-like TXT streaming parser.
- Normalize timestamps.
- Handle multi-line messages.
- Generate content hashes.
- Detect duplicate messages.
- Add fixtures and tests.

Completion criteria:

- A 1,000+ line fixture parses without loading the full file into memory.
- Expected message count and key fields pass tests.

### Phase 4. Synthetic Data Generator

Tasks:

- Implement domain catalogs.
- Implement ground-truth incident generator.
- Implement conversation renderer.
- Add controlled similarity cases.
- Add deterministic seed support.
- Write `ground-truth.json`, `factory-chat.txt`, `catalog.json`, and `expected-recall-cases.json`.

Completion criteria:

```bash
gograph generate --seed 42 --incidents 20 --messages 1000 --out ./datasets/generated
```

Same seed must produce byte-stable output.

### Phase 5. Incident Extraction

Tasks:

- Group messages into incident candidates.
- Implement rule-based baseline.
- Build Incident IR.
- Validate IR.
- Link evidence messages.
- Compare extracted incidents with ground truth.

Completion criteria:

- Report incident recall.
- Report asset linking accuracy.
- Report order linking accuracy.
- Report evidence precision.
- Report temporal accuracy.

### Phase 6. Search And Time Analysis

Tasks:

- Implement full-text baseline.
- Generate/store incident embeddings.
- Query Neo4j vector index.
- Combine keyword, vector, graph, and time signals.
- Compute deterministic timing metrics in Go.
- Return evidence-grounded results.

Completion criteria:

- Expected recall cases rank relevant incidents above keyword-only distractors.
- Results include score breakdown and evidence messages.

### Phase 7. SSH Product Interface

Tasks:

- Implement Wish-based SSH server.
- Implement password authentication from env vars.
- Implement command parser.
- Add commands: `status`, `incidents`, `open`, `recall`, `graph`, `explain`.
- Add graceful shutdown.
- Add terminal width handling.

Completion criteria:

```bash
ssh -p 2222 operator@localhost
```

The user can operate the MVP through SSH.

### Phase 8. Quality And Security

Tasks:

- Unit tests.
- Integration tests.
- Parser fuzz test.
- Race detector.
- Static analysis.
- Input size limits.
- Message length limits.
- Query timeouts.
- Context cancellation.
- SSH failed-attempt delay.
- Secret-safe logging.
- Backup and restore documentation.

Validation commands:

```bash
gofmt -w .
goimports -w .
go vet ./...
staticcheck ./...
go test ./...
go test -race ./...
```

### Phase 9. Container Image

Tasks:

- Add multi-stage Dockerfile.
- Build static Go binary.
- Run as non-root.
- Keep image tag versions fixed.
- Add `.dockerignore`.

Completion criteria:

```bash
docker build -t go-with-neo4j:local .
docker run --rm go-with-neo4j:local
```

### Phase 10. AWS EC2 Deployment

Tasks:

- Add `.fordeploy/compose.aws-demo.yaml`.
- Add `.fordeploy/deploy.aws-demo.sh`.
- Follow the deployment pattern used by the sibling `phpsucks` repository:
  local clean clone, local Docker build, `docker save`, transfer to bastion,
  transfer to private host, `docker load`, then `docker compose up -d`.
- Keep Neo4j private to Docker network.
- Expose only goGraph SSH app port.
- Use EBS-backed Neo4j data directory.
- Document backup and restore.
- Add ECR/GitHub Actions only after local MVP is stable.

Do not create AWS resources without explicit user approval.

Deployment script requirements:

- Run from WSL.
- Build from a clean local clone, not from the active working tree.
- Use clean clone root:

```bash
~/deploy-remote-repo
```

- Clone or refresh:

```bash
git@github.com:HCHJEONG/go-with-neo4j.git
```

- Reset only the clean clone to the target origin branch.
- Build a Docker image locally from the clean clone.
- Use an explicit image tag such as `go-with-neo4j:awsYYYYMMDDHHMMSS`.
- Save the image with `docker save` into a temporary tar file.
- Copy the image tar and Compose file to the bastion host.
- Copy both from bastion to the private target host.
- On the private host, require the env file to already exist.
- On the private host, copy the Compose file into the app runtime directory.
- On the private host, run `docker load`.
- On the private host, run `docker compose --env-file ... -p go-with-neo4j -f compose.aws-demo.yaml up -d`.
- Verify app SSH port readiness and container status.
- Remove temporary image tar and staged Compose files from local, bastion, and private host after success.

For ordinary app deployment, do not remove Neo4j volumes or data directories.
Do not run `docker compose rm --volumes` against the Neo4j service. If the app
container must be recreated, target only the app service unless the user
explicitly approves database maintenance.

## 15. Security Requirements

- Do not commit `.env`.
- Do not hardcode SSH password.
- Do not log passwords.
- Do not expose Neo4j ports publicly.
- Do not concatenate Cypher strings with user input.
- Use Cypher parameters.
- Do not execute LLM-generated Cypher.
- Limit file size.
- Limit message length.
- Add command timeout.
- Run production container as non-root.
- Do not ingest real private company data.

## 16. MVP Completion Criteria

MVP is complete when:

1. Go app runs directly in WSL.
2. Local Neo4j runs directly in WSL.
3. The app initializes graph schema and indexes.
4. The app ingests 1,000+ synthetic factory chat messages.
5. Repeated ingest does not duplicate messages.
6. Incidents, assets, orders, actions, and evidence are stored in Neo4j.
7. Neo4j vector index participates in similar incident search.
8. Search results include source evidence.
9. SSH password login works.
10. Core commands work from SSH.
11. Unit and integration tests pass.
12. AWS Docker Compose deployment can run Go app and Neo4j together.
13. Docker image builds.
14. README lets another developer reproduce local setup.

## 17. First Codex Implementation Request

Use this as the first implementation instruction:

```text
Read PLAN.md completely.
Implement Phase 0 and Phase 1 only.

Use Go module path github.com/HCHJEONG/go-with-neo4j.
Do not implement SSH, parser, synthetic data generation, graph schema, vector index, Docker Compose deployment, or AWS deployment yet.

Verify the current WSL, Go, Git, Java, Neo4j, and cypher-shell environment.
Create the basic repository structure.
Create .gitignore, .env.example, minimal Go app, config loading, local Neo4j readiness check, basic config tests, and README local setup instructions.

Before installing system packages, creating AWS resources, pushing to remote, or doing anything that may cost money, stop and ask for approval.
```

## 18. Final Direction

The core of goGraph is not Go, Neo4j, SSH, or AI by themselves.

The core is reconstructing this chain from fragmented shop-floor conversation:

```text
What happened
-> what asset/order/product was affected
-> who responded and when
-> what action was taken
-> what changed afterward
-> what past situation is similar
-> which original messages support that answer
```

Every implementation decision should make that chain more reliable, testable, and explainable.
