package selection

type Filter interface {
	And(other Filter) Filter
	Or(other Filter) Filter
	Xor(other Filter) Filter
	Not() Filter
}

func filterAnd(f1 Filter, f2 Filter) Filter {
	return &FilterBinaryNode{
		left:  f1,
		op:    opAnd,
		right: f2,
	}
}

func filterOr(f1 Filter, f2 Filter) Filter {
	return &FilterBinaryNode{
		left:  f1,
		op:    opOr,
		right: f2,
	}
}

func filterXor(f1 Filter, f2 Filter) Filter {
	return &FilterBinaryNode{
		left:  f1,
		op:    opXor,
		right: f2,
	}
}

func filterNot(f1 Filter) Filter {
	return &FilterUnaryNode{
		op:      opNot,
		operand: f1,
	}
}
