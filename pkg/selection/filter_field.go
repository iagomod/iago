package selection

import "github.com/kmirzavaziri/iago/pkg/model"

type FieldFilter struct {
	field FieldDescriptor
	op    FieldFilterOperator
	v     model.Value
}

type FieldFilterOperator string

const (
	Grater           = "grater"
	NotGrater        = "not_grater"
	GraterEqual      = "grater_equal"
	NotGraterEqual   = "not_grater_equal"
	Less             = "less"
	NotLess          = "not_less"
	LessEqual        = "less_equal"
	NotLessEqual     = "not_less_equal"
	Equal            = "equal"
	NotEqual         = "not_equal"
	NullSafeEqual    = "null_safe_equal"
	NullSafeNotEqual = "null_safe_not_equal"
	Between          = "between"
	NotBetween       = "not_between"
	In               = "in"
	NotIn            = "not_in"
	Is               = "is"
	NotIs            = "not_is"
	IsNull           = "is_null"
	IsNotNull        = "is_not_null"
	Like             = "like"
	NotLike          = "not_like"
	Regexp           = "regexp"
	NotRegexp        = "not_regexp"
	Has              = "has"
	NotHas           = "not_has"
)

func (f *FieldFilter) And(other Filter) Filter {
	return filterAnd(f, other)
}

func (f *FieldFilter) Or(other Filter) Filter {
	return filterOr(f, other)
}

func (f *FieldFilter) Xor(other Filter) Filter {
	return filterXor(f, other)
}

func (f *FieldFilter) Not() Filter {
	return filterNot(f)
}
