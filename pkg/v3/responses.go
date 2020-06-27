package v3

// An object to hold reusable 'Responses Object', which is provided as Node or Reference ($ref) to a 'Responses Object' definition
type Responses struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      ResponsesNode
}

// Represents 'Responses Object' as defined by the http://spec.openapis.org/oas/v3.0.3#responses-object
type ResponsesNode struct {
	Default     *Response            `json:"default,omitempty" yaml:"default,omitempty"`
	StatusCodes map[string]*Response `json:"-" yaml:"-"` // Patterned Fields; The field name MUST begin with pattern matching to ^[1-5](?:\d{2}|XX)$ , for example, 200 , 204 , 404.
	Extensions  *Extension           `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
