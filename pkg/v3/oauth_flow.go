package v3

type OAuthFlow struct {
	AuthorizationUrl string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"` // Required, applies to oauth2 ("implicit", "authorizationCode") flow
	TokenUrl         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`                 // Required, applies to oauth2 ("password", "clientCredentials", "authorizationCode")
	RefreshUrl       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"` // Required, applies to oauth2
	Extensions       *Extension        `json:"-" yaml:"-"`                               // The field name MUST begin with x-, for example, x-internal-id.
}
