package v3

type Schema struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference to a definition
	Node      SchemaNode
}

type SchemaNode struct {
	Title            string  `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf       float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool    `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
}
