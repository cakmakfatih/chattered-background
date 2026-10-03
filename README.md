# Chattered Background

Redis-backed background task worker for Chattered, implemented with Asynq.

## Run with the Chattered development stack

From the Chattered repository root, start this service and its dependencies with:

```sh
docker compose -f docker/compose.dev.yml up --build background
```

The worker has no network listener or published host port. It connects to the `redis` Compose service using `REDIS_ADDR` and `REDIS_PASSWORD` from the root `.env.dev` file.

## Task types

The worker currently handles `background:ping`, a small no-side-effect task for checking the queue path. Add production task types and handlers under `pkg/tasks` and `internal/worker` as the service's responsibilities are defined.

Tasks are delivered at least once, so handlers that perform side effects should be idempotent.

## Local development

Set `REDIS_ADDR` and `REDIS_PASSWORD` in the environment, then run:

```sh
go run ./cmd/background
```

For live reload, install [Air](https://github.com/air-verse/air) and run it from this directory:

```sh
go install github.com/air-verse/air@latest
air -c .air.toml
```

Air rebuilds the worker into `build/` whenever Go source files change. The build output is ignored by Git.

The task producer can use the shared task type and payload from `pkg/tasks` with an Asynq client configured with the same Redis address and password.
