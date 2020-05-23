package v3

type Schema struct {
	Openapi string    `json:"openapi" yaml:"openapi"` // Required
	Info    *Info     `json:"info" yaml:"info"`       // Required
	Servers []*Server `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths   *Path     `json:"paths" yaml:"paths"` // Required
}
