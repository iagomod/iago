package selection

type All struct{}

func (f *All) And(other Filter) Filter {
	return filterAnd(f, other)
}

func (f *All) Or(other Filter) Filter {
	return filterOr(f, other)
}

func (f *All) Xor(other Filter) Filter {
	return filterXor(f, other)
}

func (f *All) Not() Filter {
	return filterNot(f)
}
