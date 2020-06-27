package v3

type Components struct {
	Schemas    map[string]*Schema   `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses  map[string]*Response `json:"responses,omitempty" yaml:"responses,omitempty"`
	Extensions *Extension           `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
