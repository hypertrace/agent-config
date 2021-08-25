.PHONY: all
all: # Generates loaders for config proto.
	@echo "Generating the loaders"
	@go get ./...
	@go run cmd/generator/main.go -o $(OUT_DIR) $(SRC_DIR)
	@echo "Done."
