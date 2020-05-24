package v3

type Info struct {
	Title          string                 `json:"title" yaml:"title"` // Required
	Description    string                 `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string                 `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact               `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License               `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string                 `json:"version" yaml:"version"` // Required
	Extensions     map[string]interface{} `json:"-" yaml:"-"`             // Allows extensions to the OpenAPI Schema. The field name MUST begin with x-, for example, x-internal-id. The value can be null, a primitive, an array or an object. Can have any valid JSON format value.
}
