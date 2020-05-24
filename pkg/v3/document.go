package v3

type Document struct {
	OpenApi      string                 `json:"openapi" yaml:"openapi"` // Required
	Info         *Info                  `json:"info" yaml:"info"`       // Required
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        *Path                  `json:"paths" yaml:"paths"` // Required
	Components   *Component             `json:"components,omitempty" yaml:"components,omitempty"`
	Security     *SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []*Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Extensions   map[string]interface{} `json:"-" yaml:"-"` // Allows extensions to the OpenAPI Schema. The field name MUST begin with x-, for example, x-internal-id. The value can be null, a primitive, an array or an object. Can have any valid JSON format value.
}
