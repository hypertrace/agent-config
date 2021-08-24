package v1

import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Bool wraps the scalar value to be used in the AgentConfig object
func Bool(val bool) *wrapperspb.BoolValue {
	return &wrapperspb.BoolValue{Value: val}
}

// String wraps the scalar value to be used in the AgentConfig object
func String(val string) *wrapperspb.StringValue {
	return &wrapperspb.StringValue{Value: val}
}

func Int32(val int32) *wrapperspb.Int32Value {
	return &wrapperspb.Int32Value{Value: val}
}
