build:
	go build -o api main.go
server:
	go run main.go
test:
	go test ./... -coverprofile=coverage.out
coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($$3) > 80) != 1 )}'
report:
	go tool cover -html=coverage.out -o cover.html
check-format:
	test -z $$(go fmt ./...)
vet
	go vet ./...

.PHONY: build server test coverage report check-format vet