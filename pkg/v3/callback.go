package v3

// An object to hold reusable 'Callback Object', which is provided as Node or Reference ($ref) to a 'Callback Object' definition
type Callback struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      CallbackNode
}

// Represents 'Callback Object' as defined by http://spec.openapis.org/oas/v3.0.3#callback-object
type CallbackNode struct {
	Expression map[string]*PathItem `json:"-" yaml:"-"` // The expession key that identifies the Path Item Object is a runtime expression that can be evaluated in the context of a runtime HTTP request/response to identify the URL to be used for the callback request. Please read expression definition at http://spec.openapis.org/oas/v3.0.3#key-expression
	Extensions *Extension           `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
