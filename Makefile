GOHOSTOS:=$(shell go env GOHOSTOS)
API_TOOL_VERSION?=0.0.3
GO:=go

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find . -name *.proto -not -path './third_party/*' -not -path './.git/*'")
else
	API_PROTO_FILES=$(shell find . -name *.proto -not -path './third_party/*' -not -path './.git/*')
endif

.PHONY: tools
# install tools
tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: upgrade
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

.PHONY: api_gen
api_gen:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--validate_out=paths=source_relative,lang=go:. \
		--openapi_out=fq_schema_naming=true,default_response=false,enum_type=string:. \
		--go-errors_out=paths=source_relative:. \
		$(API_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	docker run --rm -v $(PWD):/workspace europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:$(API_TOOL_VERSION) make api_gen

.PHONY: build_api_tool
build_api_tool:
	docker build -t europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:$(API_TOOL_VERSION) -f Dockerfile.meepo-api-tool .
	docker tag europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:$(API_TOOL_VERSION) europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:latest

.PHONY: push_api_tool
push_api_tool:
	docker push europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:$(API_TOOL_VERSION)
	docker push europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:latest

.PHONY: build_and_push_api_tool
build_and_push_api_tool:
	make build_api_tool
	make push_api_tool
