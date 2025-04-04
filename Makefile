GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find . -name *.proto -not -path './third_party/*'")
else
	API_PROTO_FILES=$(shell find . -name *.proto -not -path './third_party/*')
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
	go install github.com/google/wire/cmd/wire@latest

.PHONY: api
# generate api proto
api:
	docker run --rm -v $(PWD):/workspace europe-west1-docker.pkg.dev/robust-metrics-445612-t0/meepo-api-tool/meepo-api-tool:0.0.1 protoc --proto_path=. \
	   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   --go-http_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
	   --validate_out=paths=source_relative,lang=go:. \
	   --openapi_out=fq_schema_naming=true,default_response=false:. \
	   $(API_PROTO_FILES)
