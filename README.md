# Spotify Artist API

Neste vídeo recriamos o endpoint de visualizar os detalhes de um artista.

A referência pode ser encontrada no link: 
https://developer.spotify.com/documentation/web-api/reference/get-an-artist

## YouTube

https://www.youtube.com/watch?v=xdwYp1m1GqM

## Setup

### Docker

```sh
docker compose up -d
```

### Environment

```sh
export DB_DSN='postgres://root:4321@localhost/spotify?sslmode=disable'
export APP_PORT=':8080'
```

### Migrations

```sh
migrate -path=./migrations -database=$DB_DSN up
```

### Iniciar API

```sh
go run ./cmd/server
```
