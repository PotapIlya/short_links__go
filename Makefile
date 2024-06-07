dev:
	go run ./cmd

build:
	go build ./cmd/main.go

#build-image:
#	go build ./cmd

swagger:
	swag init -g ./cmd/main.go
