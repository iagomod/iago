package model

type Value interface {
	isValue()
}

type Null struct{}

func (v Null) isValue() {}

type Number float64

func (v Number) isValue() {}

type String string

func (v String) isValue() {}

type Bool bool

func (v Bool) isValue() {}

type Map map[string]Value

func (v Map) isValue() {}

type List []Value

func (v List) isValue() {}
