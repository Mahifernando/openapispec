package v3

type Contact struct {
	Name       string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Url        string                 `json:"url,omitempty" yaml:"url,omitempty"`
	Email      string                 `json:"email,omitempty" yaml:"email,omitempty"`
	Extensions map[string]interface{} `json:"-" yaml:"-"` // Allows extensions to the OpenAPI Schema. The field name MUST begin with x-, for example, x-internal-id. The value can be null, a primitive, an array or an object. Can have any valid JSON format value.
}
