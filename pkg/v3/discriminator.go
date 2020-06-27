package v3

type Discriminator struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"` // Required
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}
