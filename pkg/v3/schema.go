package v3

// An object to hold reusable 'Schema Object', which is provided as Node or Reference ($ref) to a 'Schema Object' definition
type Schema struct {
	Reference string `json:"$ref" yaml:"$ref"` // Required if reference ($ref) to a definition
	Node      SchemaNode
}

// Represents 'Schema Object' as defined by the http://spec.openapis.org/oas/v3.0.3#schema-object
type SchemaNode struct {
	Title                string                 `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf           float64                `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum              float64                `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum              float64                `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength            uint64                 `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength            uint64                 `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern              string                 `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems             uint64                 `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems             uint64                 `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems          bool                   `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties        uint64                 `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties        uint64                 `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required             []string               `json:"required,omitempty" yaml:"required,omitempty"`
	Enum                 []interface{}          `json:"enum,omitempty" yaml:"enum,omitempty"`
	Type                 []string               `json:"type,omitempty" yaml:"type,omitempty"`
	Not                  *Schema                `json:"not,omitempty" yaml:"not,omitempty"`
	AllOf                []*Schema              `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []*Schema              `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []*Schema              `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Items                *Schema                `json:"items,omitempty" yaml:"items,omitempty"`
	Properties           map[string]*Schema     `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties *Schema                `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Description          string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Format               string                 `json:"format,omitempty" yaml:"format,omitempty"`
	Default              interface{}            `json:"default,omitempty" yaml:"default,omitempty"`
	Nullable             bool                   `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Discriminator        *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly             bool                   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly            bool                   `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	Example              interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Deprecated           bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Xml                  *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	Extensions           *Extension             `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
