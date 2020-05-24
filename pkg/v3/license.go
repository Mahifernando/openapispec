package v3

type License struct {
	Name       string     `json:"name" yaml:"name"` // Required
	Url        string     `json:"url,omitempty" yaml:"url,omitempty"`
	Extensions *Extension `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
