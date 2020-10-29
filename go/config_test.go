package config

import (
	"os"
	"testing"
)

func TestSourcesPrecedence(t *testing.T) {
	// defines the config file path
	os.Setenv("HT_CONFIG_FILE", "./testdata/config.json")

	// defines the DataCapture.HTTPHeaders.Request = true
	os.Setenv("HT_DATA_CAPTURE_HTTP_HEADERS_REQUEST", "true")

	// defines the DataCapture.HTTPHeaders.Request = true
	os.Setenv("HT_DATA_CAPTURE_HTTP_HEADERS_RESPONSE", "true")

	// loads the config
	cfg := Load()

	// use defaults
	if want, have := false, cfg.GetDataCapture().GetHTTPBody().GetResponse(); want != have {
		t.Errorf("unexpected value for dataCapture.httpBody.response, want %q, have %q", want, have)
	}

	// config file take precende over defaults
	if want, have := "api.traceable.ai", cfg.GetReporting().GetAddress(); want != have {
		t.Errorf("unexpected value for reporting.address, want %q, have %q", want, have)
	}

	// env vars take precendence over config file
	if want, have := true, cfg.GetDataCapture().GetHTTPHeaders().GetRequest(); want != have {
		t.Errorf("unexpected value for dataCapture.httpHeaders.request, want %q, have %q", want, have)
	}

	// static value take precedence over config files
	cfg.DataCapture.HTTPHeaders.Response = BoolVal(false)
}
