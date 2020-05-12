.PHONY: run
run:
	go run ./cmd/shelby/main.go

.PHONY: test
test:
	go test ./...
