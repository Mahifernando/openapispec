package v3

type Tag struct {
	Name         string                 `json:"name" yaml:"name"` // Required
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Extensions   *Extension             `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
