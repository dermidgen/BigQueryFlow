.PHONY: docs license fix vet fmt lint test build tidy
GOBIN := $(shell go env GOPATH)/bin

build:
	go build -o $(GOBIN)/bqflow -v .

all: vet lint test build buildall tidy

buildall:
	GOOS=windows go build -o /dev/null
	GOOS=linux go build -o /dev/null
	GOOS=darwin go build -o /dev/null

lint:
	(which golangci-lint || go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.22.2)
	$(GOBIN)/golangci-lint run ./...

test:
	go test -cover ./...

vet:
	go vet ./...

docker:
	docker build .

tidy:
	go mod tidy

