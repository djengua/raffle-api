build:
	go build -o api main.go
server:
	go run main.go

.PHONY: build server