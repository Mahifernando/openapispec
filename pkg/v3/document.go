package v3

// This is the root document object of the OpenAPI document as defined by http://spec.openapis.org/oas/v3.0.3#openapi-object
type Document struct {
	OpenApi      string                 `json:"openapi" yaml:"openapi"` // Required
	Info         *Info                  `json:"info" yaml:"info"`       // Required
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        *Path                  `json:"paths" yaml:"paths"` // Required
	Components   *Component             `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []*Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Extensions   *Extension             `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
