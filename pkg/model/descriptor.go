package model

type Descriptor interface {
	Table() string
	Fields() []*FieldDescriptor
	Instantiate(map[string]Value) (Model, error)
}
