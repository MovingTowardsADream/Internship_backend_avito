include .env.example
export

app run:
	go run cmd/app/main.go --config=./configs/config.yaml
.PHONY: app run

migration create:
	go run ./cmd/migrator --migrations-path=./schema
.PHONY: migration create