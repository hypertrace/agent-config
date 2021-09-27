// Code generated by github.com/hypertrace/agent-config/tools/go-generator. DO NOT EDIT.

package v1

import (
	"log"

	"google.golang.org/protobuf/proto"
)

type opts struct {
	prefix        string
	defaultConfig *AgentConfig
}

var defaultOptions = opts {
	prefix:        "HT_",
	defaultConfig: &AgentConfig{},
}

type LoadOption func(o *opts)

func WithEnvPrefix(prefix string) LoadOption {
	return func(o *opts) {
		o.prefix = prefix
	}
}

func WithDefaults(defaults *AgentConfig) LoadOption {
	return func(o *opts) {
		// The reason why we clone the message instead of reusing the one passed by the user
		// is because user might decide to change values in runtime and that is undesirable
		// without a proper API.
		var ok bool
		o.defaultConfig, ok = proto.Clone(defaults).(*AgentConfig)
		if !ok {
			log.Fatal("failed to initialize config.")
		}
	}
}
