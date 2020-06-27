package v3

type RequestBody struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType `json:"content" yaml:"content"` // Required
	Required    bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	Extensions  *Extension            `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
