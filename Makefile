GO := go
GOFMT := gofmt
LINTER := golangci-lint

SRC_DIR := pkg/nimbus

.PHONY: all
all: lint

.PHONY: lint
lint:
	$(LINTER) run

.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(SRC_DIR)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: clean
clean:
	$(GO) clean