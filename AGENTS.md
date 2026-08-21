# Project Notes

This repository is for goGraph, a backend-first Go product using Neo4j and
an SSH product interface.

The first meaningful product surface is not a web GUI. Users should eventually
connect to the application through SSH, ingest or inspect factory operation
conversation data, and recall similar incidents with source-message evidence.

The system must stay evidence-grounded. LLM output, if added later, must produce
validated intermediate data only. It must never directly generate or execute
Cypher.

## Source Of Truth

- Read `PLAN.md` before implementation work.
- Implement one phase at a time.
- Do not pull parser, SSH, vector search, or deployment work into Phase 1.
- Keep `github.com/HCHJEONG/go-with-neo4j` as the Go module path.

## Product Direction

- Backend is the product.
- SSH is the MVP user interface.
- The initial SSH interface is line-oriented, not a full-screen TUI.
- Use password authentication for the MVP, with credentials loaded from
  environment variables.
- Neo4j is the primary persistence and graph query engine.
- Neo4j vector index is part of the planned search path.
- Synthetic factory conversation data must be generated from ground truth first,
  not invented as unstructured text first.
- Search results must distinguish source facts from derived inference.
- Every important claim must reference source message IDs.

## Development Rules

- Prefer simple, explicit Go code.
- Keep runtime configuration external to the image.
- Use environment variables for configuration.
- Do not commit `.env`, credentials, SSH keys, real private chats, or production
  data.
- Do not use real company conversations or personal chats as fixtures.
- Do not concatenate Cypher with user input. Use parameters.
- Do not add broad abstractions before there are multiple real implementations.
- Do not introduce a web frontend, Kubernetes, ECS, or CI/CD deployment before
  the MVP explicitly requires it.
- Do not install system packages, create AWS resources, push to the remote, or
  trigger cost-incurring work without user approval.
- Initial setup and deployment are manual-first. Prefer documented manual
  commands before scripts, and add scripts only after the manual process is
  proven.

## Docker And Compose Direction

- Local development runs Go and Neo4j directly in WSL.
- Docker is still used locally for deployment packaging: a clean local clone is
  built into a Docker image, then saved as a tar archive.
- Remote runtime uses Docker and Docker Compose for both the Go app and Neo4j.
- Runtime hosts should not build or compile the Go application directly.
- Docker images must contain the built application artifact and runtime only,
  not local development files or secrets.
- Compose files must read secrets and environment-specific values from env
  files on the target host.
- Neo4j data must live in persistent volumes or host-mounted directories.
- Application and Neo4j containers must have separate lifecycles.
- App redeploy scripts must not remove, reset, or recreate Neo4j data unless the
  script is explicitly dedicated to database administration and the user has
  approved it.

## Deployment Rules

Deployment methodology follows the conservative pattern from `phpsucks/.fordeploy/deploy.sh`:

- Deployment is manual only unless the user explicitly requests automation.
- Deployment must be performed with shell scripts and Docker Compose.
- Before adding deployment scripts, document and execute the equivalent manual
  commands first.
- Build Docker images locally in WSL from a clean clone, not from the active
  working tree.
- Use a separate local build clone root, following the existing convention:

```bash
~/deploy-remote-repo
```

- A deployment script may clone or refresh this clean build checkout from:

```bash
git@github.com:HCHJEONG/go-with-neo4j.git
```

- The script should checkout the target branch from origin, hard reset the clean
  clone, and clean untracked files there. Do not do this in the user's active
  working tree.
- Build the Docker image locally from the clean clone.
- Tag images with an explicit AWS timestamp or commit-derived tag, for example
  `go-with-neo4j:awsYYYYMMDDHHMMSS`. Do not rely only on `latest`.
- Save the built image to a `.tar` archive locally.
- Transfer the image tarball and target Compose file to the bastion host when a
  bastion is used.
- From the bastion host, transfer the image tarball and target Compose file to
  the private target host.
- Load the image on the private target host with Docker.
- Copy the transferred Compose file into the app runtime directory.
- Run or update the service on the private target host with `docker compose`.
- Verify the running service after deployment.
- Clean temporary tarballs and transferred Compose staging files after success.

## Target Layout

Use target-specific runtime paths. For `aws-demo`, prefer the existing
home-based layout:

```text
/home/ubuntu/go-with-neo4j
/home/ubuntu/docker_images/go-with-neo4j
```

Recommended runtime files:

```text
/home/ubuntu/go-with-neo4j/.env
/home/ubuntu/go-with-neo4j/compose.aws-demo.yaml
/home/ubuntu/go-with-neo4j/neo4j/data
/home/ubuntu/go-with-neo4j/neo4j/logs
```

Rules:

- Deployment scripts may create directories.
- Deployment scripts must not overwrite existing env files containing secrets.
- The target env file must exist before production-like deployment.
- Neo4j Bolt and browser ports must not be exposed to the public internet.
- Only the goGraph application SSH port should be exposed, and only as
  deliberately configured.

## AWS Demo Pattern

The current sibling deployment pattern uses:

- a local WSL script,
- a clean local clone,
- local Docker image build,
- `docker save`,
- transfer to a bastion host,
- second transfer from bastion to private host,
- target env file pre-existence check,
- `sudo docker load` on the private host,
- final deployment on a private host,
- `docker compose --env-file ... up -d`.

If this repository adds `.fordeploy`, keep scripts target-specific, for example:

```text
.fordeploy/
├── compose.aws-demo.yaml
└── deploy.aws-demo.sh
```

The Compose file should define separate services for:

- `app`
- `neo4j`

The deployment script should replace only the app container for ordinary app
deployments. It must preserve Neo4j volumes and data directories.

Unlike the WordPress deployment in `phpsucks`, goGraph must not use
`docker compose rm --volumes` for the Neo4j service during normal app deploys.
If the app container needs to be recreated, target only the app service and keep
Neo4j running unless the user explicitly approves database maintenance.

## Secrets And Runtime Files

- Never include env files in Docker images.
- Never include credentials in Docker images.
- Never commit SSH passwords, host keys, private keys, or Neo4j passwords.
- Do not print secrets in logs.
- `aws-demo` setup scripts may create template files but must not generate or
  overwrite production-like secrets.
- Development-only scripts may generate demo credentials only if they preserve
  existing env files.

## Validation

Before claiming local development readiness, verify:

```bash
go version
neo4j --version
cypher-shell --version
go test ./...
```

Before claiming deployment packaging readiness, verify from the clean build
clone:

```bash
docker build -t go-with-neo4j:local <clean-clone-path>
docker save -o /tmp/go-with-neo4j-image.tar go-with-neo4j:local
```

For remote deployment scripts, verify on the target host:

```bash
docker compose --env-file .env -p go-with-neo4j -f compose.aws-demo.yaml ps
docker compose --env-file .env -p go-with-neo4j -f compose.aws-demo.yaml logs --tail 200
```

Application health verification should eventually check:

- app container is running,
- Neo4j container is healthy,
- the app can connect to Neo4j,
- the SSH application port accepts a connection,
- existing Neo4j data was preserved.

## Agent Guidance

- Read `PLAN.md` before broad changes.
- Keep Phase 1 limited to repository setup, config, local Neo4j readiness, and
  app readiness.
- When adding deployment scripts, follow the local-build, tar-transfer,
  remote-compose pattern from sibling repositories.
- Do not add registry-based deployment, GitHub Actions deployment, ECS, or
  Kubernetes unless explicitly requested.
- Do not stop, remove, or recreate the Neo4j container or volume from an app
  deployment script unless that script is clearly marked as database
  administration and the user has approved it.
- Make paths, image names, container names, ports, env files, and target hosts
  easy to inspect before execution.
- Clearly encode the target in script names or config files to avoid accidental
  cross-environment deployment.
