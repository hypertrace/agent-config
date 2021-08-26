// Code generated by github.com/hypertrace/agent-config/tools/go-generator. DO NOT EDIT.

package v1

type opts struct {
	prefix        string
	defaultConfig AgentConfig
}

var defaultOptions = opts{
	prefix:        "HT_",
	defaultConfig: AgentConfig{},
}

type LoadOption func(o *opts)

func WithEnvPrefix(prefix string) LoadOption {
	return func(o *opts) {
		o.prefix = prefix
	}
}

func WithDefaults(defaults AgentConfig) LoadOption {
	return func(o *opts) {
		o.defaultConfig = defaults
	}
}
