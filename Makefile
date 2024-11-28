GO := go
GOFMT := gofmt
LINTER := golangci-lint

SRC_DIR := pkg/nimbus

.PHONY: all
all: lint

.PHONY: lint
lint:
	$(LINTER) run -v

.PHONY: lint-fix
lint-fix:
	$(LINTER) run -v --fix

.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(SRC_DIR)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: clean
clean:
	$(GO) clean