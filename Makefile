generate-go:
	docker build -t agent-config-go -f go/Dockerfile .
	docker run \
	-v $(PWD)/go:/go/src/agent-config/lib \
	agent-config-go \
	run lib/cmd/generator/main.go -o ./lib config.proto

test-go:
	docker build -t agent-config-go .
	@cd go; go test ./...

generate: generate-go
