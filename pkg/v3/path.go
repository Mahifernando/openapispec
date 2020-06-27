package v3

type Path struct {
	PathItems  map[string]*PathItem `json:"-" yaml:"-"` // A relative path to an individual endpoint. The field name MUST begin with a forward slash (/) , for example, /pets/{petId} /pets/mine.
	Extensions *Extension           `json:"-" yaml:"-"` // The field name MUST begin with x-, for example, x-internal-id.
}
