# Golang API (Gin Gonic)

## Setup

1. Copy .env.template and rename to .env, then set environment variables
2. Then run `docker compose up` to run database
3. Run `go run server.go` to run server (You can also build executable before running server)

## API Specification

**POST** /api/strong_password_steps

**request: json**

`{
"init_password": "password_string"
}`

**response: json**

`{
"num_of_steps": 0
}`

## To run unit test

`go test ./services`

## To use API with cURL

`curl -X POST -H "Content-Type: application/json" -d '{
"init_password": "password"
}' YOUR_HOST:YOUR_PORT/api/strong_password_steps`
