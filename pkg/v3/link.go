package v3

type Link struct {
	OperationId  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	OperationRef string                 `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`
	Extensions   *Extension             `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
