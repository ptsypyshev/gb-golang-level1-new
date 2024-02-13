.PHONY: dependencies
dependencies:
	go mod download -x

.PHONY: build
build:
	go build cmd/main.go

.PHONY: run
run:
	go run cmd/main.go -load cmd/data.json

.PHONY: test
test:
	go test ./...

.PHONY: coverage-html
coverage-html:
	mkdir -p tmp
	go test ./... -v -coverprofile=tmp/coverage.out
	go tool cover -html=tmp/coverage.out
