BREAKING_CHECK_AGAINST := .git\#branch=main

lint: ## Lints the proto files.
	buf lint

breaking: ## Checks the proto files for breaking changes
	buf breaking --against $(BREAKING_CHECK_AGAINST)

generate: ## generates code for all languages
	buf generate
