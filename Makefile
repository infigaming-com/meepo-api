GOHOSTOS:=$(shell go env GOHOSTOS)
GO := go

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: upgrade
# upgrade
upgrade:
	@echo "Fetching latest Go 1.23.x version..."
	$(eval LATEST_GO_VERSION := $(shell curl -s https://go.dev/dl/?mode=json | grep -o '"version": "go1\.23\.[0-9]*"' | sort -V | tail -n1 | cut -d'"' -f4 | sed 's/go//'))
	@if [ -z "$(LATEST_GO_VERSION)" ]; then \
		echo "Error: No Go 1.23.x version found"; \
		exit 1; \
	fi
	@echo "Latest Go 1.23.x version: $(LATEST_GO_VERSION)"
	@if [ "$(GOHOSTOS)" = "darwin" ]; then \
		sed -i '' 's/^go .*/go $(LATEST_GO_VERSION)/' go.mod; \
	else \
		sed -i 's/^go .*/go $(LATEST_GO_VERSION)/' go.mod; \
	fi
	@echo "Updating dependencies..."
	@$(GO) get -u ./...
	@$(GO) mod tidy
	@echo "Go version and dependencies updated successfully!"
