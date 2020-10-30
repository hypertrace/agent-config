USE_DOCKER ?= 0

generate-go:
ifeq ($(USE_DOCKER), 1)
	@docker build --no-cache -t agent-config-go -f go/Dockerfile .
	docker run \
	-v $(PWD)/go:/go/src/github.com/hypertrace/agent-config/go \
	agent-config-go \
	run ./cmd/generator/main.go -o ./config ../config.proto 
else
	@cd go; go run cmd/generator/main.go -o ./config ../config.proto
endif

test-go:
ifeq ($(USE_DOCKER), 1)
	@docker build -t agent-config-go -f go/Dockerfile .
	docker run \
	-v $(PWD)/go:/go/src/agent-config/lib \
	agent-config-go \
	test ./...
else
	@cd go; go test ./...
endif

generate: generate-go
