.PHONY: all
all: # Generates loaders for config proto.
	@echo "Generating the loaders"
	@cd cmd/generator; go get ./...
	@cd cmd/generator; go run . -out $(OUT_DIR) -opt-module github.com/hypertrace/agent-config/gen/go $(ROOT)
	@echo "Done."
