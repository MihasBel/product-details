.PHONY: mod
mod:
	go mod tidy

.PHONY: build
build:
	go build \
			--trimpath \
			-o bin/app/product-details \
			cmd/http/main.go

.PHONY: check
check:
	golangci-lint run -v --config .golangci.yml

.PHONY: test
test:
	go test -v ./...

.PHONY: race
race:
	go test -v -race ./...

.PHONY: cover
cover:
	./scripts/coverage.sh html

.PHONY: run
run: build
	go run cmd/http/main.go --config=configs/local-env/env.json

.PHONY: swag
swag:
	swag init -g cmd/http/main.go