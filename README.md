# Agent Config

Agent config contains the configuration specs for the Hypertrace agents.

## Conventions

In order to make sure compatibility across languages and versions, the following conventions must be followed:

- Fields must not be removed, if there is a need to deprecate a config settings, the field should be marked as deprecated.
- Fields should not depend on each others, every field should be independent and self contained to avoid config logic in the libraries.
- Config must not be modified in runtime by the user through code.

## Contributing

Before submitting a PR, make sure you run the linter:

```bash
protolint config.proto
```
