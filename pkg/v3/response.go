package v3

type Response struct {
	Description string             `json:"description" yaml:"description"` // Required
	Headers     map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     *MediaType         `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*Link   `json:"links,omitempty" yaml:"links,omitempty"`
	Extensions  *Extension         `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
