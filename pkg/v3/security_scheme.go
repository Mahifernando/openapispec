package v3

// An object to hold reusable 'Security Scheme Object', which is provided as Node or Reference ($ref) to a 'Security Scheme Object' definition
type SecurityScheme struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      SecuritySchemeNode
}

// Represents 'Security Scheme Object' as defined by thttp://spec.openapis.org/oas/v3.0.3#security-scheme-object
type SecuritySchemeNode struct {
	Type             string      `json:"type" yaml:"type"` // Required
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`     // Required, applies to 'apiKey' scheme
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`         // Required, applies to 'apiKey' scheme
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"` // Required, applies to 'http' scheme
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`                       // Required, applies to 'oauth2' scheme
	OpenIdConnectUrl string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"` // Required, applies to 'openIdConnect' scheme
	Extensions       *Extension  `json:"-" yaml:"-"`                                                   // The field name MUST begin with x-, for example, x-internal-id.
}
