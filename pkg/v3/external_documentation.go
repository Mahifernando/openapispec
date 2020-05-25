package v3

type ExternalDocumentation struct {
	Url         string     `json:"url" yaml:"url"` // Required
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Extensions  *Extension `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
