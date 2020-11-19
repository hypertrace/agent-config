lint:
	protolint config.proto

generate-env-vars:
	docker build -t hypertrace/agent-config/env-vars-generator tools/env-vars-generator
	docker run \
	-v $(PWD)/ENV_VARS.md:/usr/local/ENV_VARS.md \
	-v $(PWD)/config.proto:/usr/local/config.proto \
	hypertrace/agent-config/env-vars-generator \
	-o /usr/local/ENV_VARS.md \
	/usr/local/config.proto