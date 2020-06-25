package v3

// An object to hold reusable 'Header Object', which is provided as Node or Reference ($ref) to a 'Header Object' definition
type Header struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      HeaderNode
}

// Represents 'Header Object' as defined by the http://spec.openapis.org/oas/v3.0.3#header-object
type HeaderNode struct {
	Description     string              `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool                `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool                `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           string              `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         bool                `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   bool                `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          map[string]*Schema  `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content         *MediaType          `json:"content,omitempty" yaml:"content,omitempty"`
	Example         interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`
	Extensions      *Extension          `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
