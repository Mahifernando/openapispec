package v3

type Encoding struct {
	ContentType   string  `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       *Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string  `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool    `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
