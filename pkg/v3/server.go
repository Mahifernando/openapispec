package v3

type Server struct {
	Url         string                     `json:"url" yaml:"url"` // Required
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
	Extensions  *Extension                 `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
