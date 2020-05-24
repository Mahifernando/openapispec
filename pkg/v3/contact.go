package v3

type Contact struct {
	Name       string     `json:"name,omitempty" yaml:"name,omitempty"`
	Url        string     `json:"url,omitempty" yaml:"url,omitempty"`
	Email      string     `json:"email,omitempty" yaml:"email,omitempty"`
	Extensions *Extension `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
