package model

type Model interface {
	Descriptor() Descriptor
	FieldPointer(*FieldDescriptor) any
}
