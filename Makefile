.PHONY: default help lint-proto generate-env-vars init-git-submodule breaking-proto generate-proto clean

default: help

help: ## Prints this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

clean: ## Cleans all build generated artifacts
	rm -rf ./gen

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
	@# Generates pb struct
	$(MAKE) -C ./proto generate
	
	@# Generates pb loaders
	@SRC_DIR=$(PWD)/proto OUT_DIR=$(PWD)/gen/go \
	$(MAKE) -C ./tools/go-generator

	@# Runs go mod tidy for all modules
	@find $(PWD)/gen/go \( -name vendor -o -name '[._].*' -o -name node_modules \) -prune -o -name go.mod -print | sed 's:/go.mod::' | xargs -I {} bash -c 'cd {}; go mod tidy'

generate-env-vars: init-git-submodule ## Generates the ENV_VARS.md with all environment variables.
	docker build -t hypertrace/agent-config/env-vars-generator tools/env-vars-generator
	touch $(PWD)/ENV_VARS.md # makes sure this is created as a file and not as a directory
	docker run \
	-v $(PWD)/ENV_VARS.md:/usr/local/ENV_VARS.md \
	-v $(PWD)/proto/hypertrace/agent/config/v1/config.proto:/usr/local/config.proto \
	hypertrace/agent-config/env-vars-generator \
	-o /usr/local/ENV_VARS.md \
	/usr/local/config.proto

init-git-submodule:
	git submodule update --init
