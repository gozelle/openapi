package openapi3filter

import "github.com/gozelle/openapi/openapi3"

// DefaultOptions do not set an AuthenticationFunc.
// A spec with security schemes defined will not pass validation
// unless an AuthenticationFunc is defined.
var DefaultOptions = &Options{}

// Options used by ValidateRequest and ValidateResponse
type Options struct {
	// Set ExcludeRequestBody so ValidateRequest skips request body validation
	ExcludeRequestBody bool
	
	// Set ExcludeResponseBody so ValidateResponse skips response body validation
	ExcludeResponseBody bool
	
	// Set IncludeResponseStatus so ValidateResponse fails on response
	// status not defined in OpenAPI spec
	IncludeResponseStatus bool
	
	MultiError bool
	
	// See NoopAuthenticationFunc
	AuthenticationFunc AuthenticationFunc
	
	// Indicates whether default values are set in the
	// request. If true, then they are not set
	SkipSettingDefaults bool
	
	customSchemaErrorFunc CustomSchemaErrorFunc
}

// CustomSchemaErrorFunc allows for custom the schema error message.
type CustomSchemaErrorFunc func(err *openapi3.SchemaError) string

// WithCustomSchemaErrorFunc sets a function to override the schema error message.
// If the passed function returns an empty string, it returns to the previous Error() implementation.
func (o *Options) WithCustomSchemaErrorFunc(f CustomSchemaErrorFunc) {
	o.customSchemaErrorFunc = f
}
