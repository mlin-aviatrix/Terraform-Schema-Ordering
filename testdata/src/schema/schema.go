package schema

type Resource struct {
	Schema *Schema
}

type Schema struct {
	Type string
	Required bool
	Optional bool
	Computed bool
	Default string


	Elem string
	Description string
}