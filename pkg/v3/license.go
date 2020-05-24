package v3

type License struct {
	Name       string                 `json:"name" yaml:"name"` // Required
	Url        string                 `json:"url,omitempty" yaml:"url,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:"-"` // Allows extensions to the OpenAPI Schema. The field name MUST begin with x-, for example, x-internal-id. The value can be null, a primitive, an array or an object. Can have any valid JSON format value.
}
