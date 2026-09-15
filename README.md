# To-do backend

A Go/Gin REST API for user accounts and task management, backed by MongoDB.

## Requirements

- Go 1.19 or newer
- MongoDB

## Run locally

1. Copy `.env.example` to `.env`.
2. Set `JWT_SECRET` to a long random value.
3. Start MongoDB.
4. Run the API:

```sh
go run . start_http
```

The server listens on `APP_PORT` (default `9090`). Verify it with:

```sh
curl http://localhost:9090/ping
```

## Authentication

Create an account:

```sh
curl -X POST http://localhost:9090/to-do/v1/user/sign-up \
  -H 'Content-Type: application/json' \
  -d '{"name":"Example User","username":"example-user","password":"Example1!","phone_number":"9900923821"}'
```

Log in to receive a bearer token:

```sh
curl -X POST http://localhost:9090/to-do/v1/user/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"example-user","password":"Example1!"}'
```

All task endpoints require `Authorization: Bearer <token>`. The authenticated account is used for task ownership; request-body usernames are ignored for authorization.

## Task endpoints

- `POST /to-do/v1/task` — create a task
- `PUT /to-do/v1/task` — update a task
- `PATCH /to-do/v1/task` — change task state
- `GET /to-do/v1/tasks` — list the authenticated user’s tasks

Run the test suite with:

```sh
make test
```
