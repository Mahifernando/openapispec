package v3

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default" yaml:"default"` // Required
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}
