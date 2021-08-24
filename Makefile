.PHONY: default help lint generate-env-vars init-git-submodule breaking generate clean

default: help

help: ## Prints this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

clean: ## Cleans all build generated artifacts
	find . -type f -name '*.pb.go' -exec rm {} +

install:
	go install -mod=mod \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		github.com/bufbuild/buf/cmd/buf

lint-proto:
	$(MAKE) -C ./proto lint

breaking-proto:
	$(MAKE) -C ./proto breaking

generate-proto:
	$(MAKE) -C ./proto generate

generate-env-vars: init-git-submodule ## Generates the ENV_VARS.md with all environment variables.
	docker build -t hypertrace/agent-config/env-vars-generator tools/env-vars-generator
	touch $(PWD)/ENV_VARS.md # makes sure this is created as a file and not as a directory
	docker run \
	-v $(PWD)/ENV_VARS.md:/usr/local/ENV_VARS.md \
	-v $(PWD)/proto/org/hypertrace/agent/config/v1/config.proto:/usr/local/config.proto \
	hypertrace/agent-config/env-vars-generator \
	-o /usr/local/ENV_VARS.md \
	/usr/local/config.proto

init-git-submodule:
	git submodule update --init
