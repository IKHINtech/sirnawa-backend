generate-doc:
	swag init -g cmd/main.go

run:
	go run cmd/main.go

app-build:
	docker-compose up --build
