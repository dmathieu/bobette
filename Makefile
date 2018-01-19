GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

ci: tidy test

test:
	go test -race -v ./...
