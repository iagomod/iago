package model

type FieldDescriptor struct {
	name       string
	definition Type
}

func NewFieldDescriptor(name string, definition Type) *FieldDescriptor {
	return &FieldDescriptor{
		name:       name,
		definition: definition,
	}
}

func (d *FieldDescriptor) Name() string {
	return d.name
}

func (d *FieldDescriptor) Definition() Type {
	return d.definition
}
