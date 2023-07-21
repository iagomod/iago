package selection

type FilterUnaryNode struct {
	op      FilterUnaryOperator
	operand Filter
}

type FilterUnaryOperator string

const (
	opNot FilterUnaryOperator = "not"
)

func (f *FilterUnaryNode) And(other Filter) Filter {
	return filterAnd(f, other)
}

func (f *FilterUnaryNode) Or(other Filter) Filter {
	return filterOr(f, other)
}

func (f *FilterUnaryNode) Xor(other Filter) Filter {
	return filterXor(f, other)
}

func (f *FilterUnaryNode) Not() Filter {
	return filterNot(f)
}
