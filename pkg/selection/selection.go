package selection

import (
	"github.com/iagomod/iago/pkg/model"
)

type Selection interface {
	Errors() Errors

	Filter(filter Filter) Selection
	GetFilter() Filter

	Slice(start uint, stop uint) Selection
	SliceTillEnd(start uint) Selection
	Page(pageNumber uint, pageSize uint) Selection
	GetLimitOffset() (limit int, offset uint)

	Sort(fd FieldDescriptor, order Order) Selection
	CancelSorts() Selection
	GetSorts() []*SortClause
}

type selection struct {
	d      model.Descriptor
	errors Errors

	filter Filter
	limit  int
	offset uint
	sorts  []*SortClause
}

func NewSelection(d model.Descriptor) Selection {
	return &selection{
		d:      d,
		errors: Errors{},
		filter: &All{},
		limit:  -1,
		offset: 0,
	}
}

func (s *selection) Errors() Errors {
	return s.errors
}

func (s *selection) Filter(filter Filter) Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	s.filter.And(filter)

	return s
}

func (s *selection) GetFilter() Filter {
	return s.filter
}

func (s *selection) Slice(start uint, stop uint) Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	if start > stop {
		return s.addError(ErrStartAfterStop)
	}

	s.offset = start
	s.limit = int(stop - start)

	return s
}

func (s *selection) SliceTillEnd(start uint) Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	s.offset = start
	s.limit = -1

	return s
}

func (s *selection) Page(pageNumber uint, pageSize uint) Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	s.offset = pageNumber * pageSize
	s.limit = int(pageSize)

	return s
}

func (s *selection) GetLimitOffset() (limit int, offset uint) {
	return s.limit, s.offset
}

func (s *selection) Sort(fd FieldDescriptor, order Order) Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	s.sorts = append(s.sorts, &SortClause{FieldName: fd.Name(), Order: order})

	return s
}

func (s *selection) CancelSorts() Selection {
	if len(s.Errors()) != 0 {
		return s
	}

	s.sorts = []*SortClause{}

	return s
}

func (s *selection) GetSorts() []*SortClause {
	return s.sorts
}

func (s *selection) addError(err error) *selection {
	s.errors = append(s.errors, err)

	return s
}
