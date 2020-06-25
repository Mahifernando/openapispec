package v3

type MediaType struct {
	Schema     map[string]*Schema  `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example    interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	Examples   map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding   *Encoding           `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	Extensions *Extension          `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
