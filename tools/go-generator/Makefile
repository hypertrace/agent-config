OPT_MODULE ?= github.com/hypertrace/agent-config/gen/go
ENV_PREFIX ?= HT_

.PHONY: all
all: # Generates loaders for config proto.
	@echo "Generating the loaders"
	@cd cmd/generator; go get ./...
	@cd cmd/generator; go run . -out $(OUT_DIR) -opt-module $(OPT_MODULE) -env-prefix $(ENV_PREFIX) $(ROOT)
	@echo "Done."
