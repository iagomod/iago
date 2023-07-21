package model

type Type interface {
	isType()
}

type Text struct {
	MaxLength uint
	Optional  bool
}

func (t Text) isType() {}
