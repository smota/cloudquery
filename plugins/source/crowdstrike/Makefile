# Generate mocks for mock/unit testing 
.PHONY: gen-mocks
gen-mocks:
	go generate ./client/...

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-docs