package v1

type opts struct {
	prefix        string
	defaultConfig AgentConfig
}

var defaultOptions = opts{
	prefix: "HT_",
	defaultConfig: AgentConfig{
		PropagationFormats: []PropagationFormat{PropagationFormat_TRACECONTEXT},
		DataCapture: &DataCapture{
			HttpHeaders: &Message{
				Request:  Bool(true),
				Response: Bool(true),
			},
			HttpBody: &Message{
				Request:  Bool(true),
				Response: Bool(true),
			},
			RpcMetadata: &Message{
				Request:  Bool(true),
				Response: Bool(true),
			},
			RpcBody: &Message{
				Request:  Bool(true),
				Response: Bool(true),
			},
			BodyMaxSizeBytes: Int32(131072),
		},
		Reporting: &Reporting{
			Endpoint:          String("http://localhost:9411/api/v2/spans"),
			Secure:            Bool(false),
			TraceReporterType: TraceReporterType_ZIPKIN,
		},
	},
}

type LoadOption func(o *opts)

func WithPrefix(prefix string) LoadOption {
	return func(o *opts) {
		o.prefix = prefix
	}
}

func WithDefaults(defaults AgentConfig) LoadOption {
	return func(o *opts) {
		o.defaultConfig = defaults
	}
}
