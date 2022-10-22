.DEFAULT_GOAL := codecheck

.PHONY: codecheck
codecheck: lint

# Requires ~ https://github.com/golangci/golangci-lint
.PHONY: lint
lint:
	~/go/bin/golangci-lint run