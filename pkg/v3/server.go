package v3

type Server struct {
	Url         string                     `json:"url" yaml:"url"` // Required
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
	Extensions  map[string]interface{}     `json:"-" yaml:"-"` // Allows extensions to the OpenAPI Schema. The field name MUST begin with x-, for example, x-internal-id. The value can be null, a primitive, an array or an object. Can have any valid JSON format value.
}
