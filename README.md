[![Build Status][ci-img]][ci]

# Agent Config

Agent config contains the configuration specs for the Hypertrace agents.

Configs can be set through config file, environment variables or code. For a list of supported environment variables have a look at [this list](ENV_VARS.md).

## Conventions

In order to make sure compatibility across languages and versions, the following conventions must be followed:

- All messages and fields should include a comment, describing what the field describes and/or the effect on a certain value, e.g. "when `false`, permits connecting to the trace endpoint without a certificate".
- Fields must not be removed, if there is a need to deprecate a config settings, the field should be marked as deprecated.
- Fields should not depend on each others, every field should be independent and self contained to avoid config logic in the libraries.
- Time duration fields should include a suffix indicating the time unit e.g. `retryIntervalSeconds`.
- Config must not be modified in runtime by the user through code.

## Contributing

Before submitting a PR, make sure you run the linter & generate the go modules:

```bash
protolint config.proto
```
```bash
make generate-proto generate-env-vars
```

[ci-img]: https://github.com/hypertrace/agent-config/workflows/lint%20protobuf/badge.svg
[ci]: https://github.com/hypertrace/agent-config/actions

Initialize the git submodules before generating the env vars.
```bash
git submodule update  --init  --recursive
```

### Updating Submodules

While updating submodules, do not add your commits inside submodule directories.
After pulling new changes in the submodule,
if you see some untracked changes in submodule, delete those untracked changes in the submodule.  
If you see some new commits in nested submodule, update the nested submodule in your local to the commit it points to in the new version of the submodule.
After these, there should not be any changes to be committed in the submodule directory.
In the root directory, add and commit the change to update submodules.