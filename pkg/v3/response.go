package v3

// An object to hold reusable 'Response Object', which is provided as Node or Reference ($ref) to a 'Response Object' definition
type Response struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      ResponseNode
}

// Represents 'Response Object' as defined by the http://spec.openapis.org/oas/v3.0.3#response-object
type ResponseNode struct {
	Description string             `json:"description" yaml:"description"` // Required
	Headers     map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     *MediaType         `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*Link   `json:"links,omitempty" yaml:"links,omitempty"`
	Extensions  *Extension         `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
