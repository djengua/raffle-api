GO_VERSION := 1.2

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz
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
vet:
	go vet ./...
install-lint:
	sudo curl -sSfL \
	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2
static-check:
	golangci-lint run

.PHONY: build server test coverage report check-format vet setup install-lint static-check