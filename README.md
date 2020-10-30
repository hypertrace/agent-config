# Agent Config

Agent config contains the configuration specs for the Hypertrace agents.

For specific documentation about language implementations:

- [Go](./go/README.md)

## Implementation

For every language, this library must provide:

- the models representing the config data
- a explicit set of default values for each programming language
- the logic for loading and parsing the config from a JSON file, env vars and a public API to set the values in code. Precedence of this values should be:
    1. code
    2. env vars
    3. config file
    4. defaults

What this library should not provide is:

- Initializers for tracing or agents.
- API for accessing the config by the user.

## Conventions

In order to make sure compatibility across languages and versions, the following conventions must be followed:

- Fields must not be removed, if there is a need to deprecate a config settings, the field should be marked as deprecated.
- Fields should not depend on each others, every field should be independent and self contained to avoid config logic in the libraries.
- Config must not be modified in runtime by the user through code.
