# Chattered Background

Redis-backed background task worker for Chattered, implemented with Asynq.

## Run with the Chattered development stack

From the Chattered repository root, start this service and its dependencies with:

```sh
docker compose -f docker/compose.dev.yml up --build background
```

The service exposes `GET /ping` on its internal HTTP port `8080` and returns plain-text `pong`. The Compose service does not publish a host port; other containers on the development network, including Caddy, can reach it at `background:8080`. The worker connects to the `redis` Compose service using `REDIS_ADDR` and `REDIS_PASSWORD` from the root `.env.dev` file.

## Task types

The worker currently handles `background:ping`, a small no-side-effect task for checking the queue path. Add production task types and handlers under `pkg/tasks` and `internal/worker` as the service's responsibilities are defined.

Tasks are delivered at least once, so handlers that perform side effects should be idempotent.

## Local development

In a Chattered checkout, local development uses the root `.env.dev` file for `REDIS_ADDR` and `REDIS_PASSWORD`. Keep that file outside this submodule and do not commit its secrets.

From this directory, run Air:

```sh
air -c .air.toml
```

Air loads `../.env.dev`, rebuilds the worker into `build/`, and restarts it when Go source files change. If you run the worker directly with `go run ./cmd/background`, provide `REDIS_ADDR` and `REDIS_PASSWORD` in the shell environment first.

Install [Air](https://github.com/air-verse/air) if it is not already available:

```sh
go install github.com/air-verse/air@latest
```

The build output is ignored by Git.

The task producer can use the shared task type and payload from `pkg/tasks` with an Asynq client configured with the same Redis address and password.
