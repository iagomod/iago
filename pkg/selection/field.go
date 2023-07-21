package selection

import (
	"github.com/iagomod/iago/pkg/model"
)

type FieldDescriptor interface {
	Name() string

	Grater(v model.Value) Filter
	NotGrater(v model.Value) Filter
	GraterEqual(v model.Value) Filter
	NotGraterEqual(v model.Value) Filter
	Less(v model.Value) Filter
	NotLess(v model.Value) Filter
	LessEqual(v model.Value) Filter
	NotLessEqual(v model.Value) Filter
	Equal(v model.Value) Filter
	NotEqual(v model.Value) Filter
	NullSafeEqual(v model.Value) Filter
	NullSafeNotEqual(v model.Value) Filter
	Between(v model.Value) Filter
	NotBetween(v model.Value) Filter
	In(v model.Value) Filter
	NotIn(v model.Value) Filter
	Is(v model.Value) Filter
	NotIs(v model.Value) Filter
	IsNull(v model.Value) Filter
	IsNotNull(v model.Value) Filter
	Like(v model.Value) Filter
	NotLike(v model.Value) Filter
	Regexp(v model.Value) Filter
	NotRegexp(v model.Value) Filter
	Has(v model.Value) Filter
	NotHas(v model.Value) Filter
}

type fieldDescriptor struct {
	name string
}

func NewFD(mfd *model.FieldDescriptor) FieldDescriptor {
	return &fieldDescriptor{
		name: mfd.Name(),
	}
}

func (fd *fieldDescriptor) Name() string {
	return fd.name
}

func (fd *fieldDescriptor) Grater(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Grater,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotGrater(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotGrater,
		v:     v,
	}
}

func (fd *fieldDescriptor) GraterEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    GraterEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotGraterEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotGraterEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) Less(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Less,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotLess(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotLess,
		v:     v,
	}
}

func (fd *fieldDescriptor) LessEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    LessEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotLessEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotLessEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) Equal(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Equal,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) NullSafeEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NullSafeEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) NullSafeNotEqual(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NullSafeNotEqual,
		v:     v,
	}
}

func (fd *fieldDescriptor) Between(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Between,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotBetween(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotBetween,
		v:     v,
	}
}

func (fd *fieldDescriptor) In(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    In,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotIn(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotIn,
		v:     v,
	}
}

func (fd *fieldDescriptor) Is(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Is,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotIs(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotIs,
		v:     v,
	}
}

func (fd *fieldDescriptor) IsNull(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    IsNull,
		v:     v,
	}
}

func (fd *fieldDescriptor) IsNotNull(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    IsNotNull,
		v:     v,
	}
}

func (fd *fieldDescriptor) Like(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Like,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotLike(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotLike,
		v:     v,
	}
}

func (fd *fieldDescriptor) Regexp(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Regexp,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotRegexp(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotRegexp,
		v:     v,
	}
}

func (fd *fieldDescriptor) Has(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    Has,
		v:     v,
	}
}

func (fd *fieldDescriptor) NotHas(v model.Value) Filter {
	return &FieldFilter{
		field: fd,
		op:    NotHas,
		v:     v,
	}
}
