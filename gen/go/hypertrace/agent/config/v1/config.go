// Code generated by github.com/hypertrace/agent-config/tools/go-generator. DO NOT EDIT.

package v1

// LoadFromEnv loads env and default values on existing AgentConfig instance
// where defaults only overrides empty values while env vars can override all
// of them.
func (x *AgentConfig) LoadFromEnv(opts ...LoadOption) {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}

	x.loadFromEnv(options.prefix, &options.defaultConfig)
}