package selection

type FilterBinaryNode struct {
	left  Filter
	op    FilterBinaryOperator
	right Filter
}

type FilterBinaryOperator string

const (
	opAnd FilterBinaryOperator = "and"
	opOr  FilterBinaryOperator = "or"
	opXor FilterBinaryOperator = "xor"
)

func (f *FilterBinaryNode) And(other Filter) Filter {
	return filterAnd(f, other)
}

func (f *FilterBinaryNode) Or(other Filter) Filter {
	return filterOr(f, other)
}

func (f *FilterBinaryNode) Xor(other Filter) Filter {
	return filterXor(f, other)
}

func (f *FilterBinaryNode) Not() Filter {
	return filterNot(f)
}
