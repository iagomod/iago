package selection

import (
	"errors"
	"strings"
)

var ErrStartAfterStop = errors.New("[Slice] start is after stop")

type Errors []error

func (ee Errors) Error() string {
	ss := make([]string, len(ee))
	for i, e := range ee {
		ss[i] = e.Error()
	}
	return strings.Join(ss, "\n")
}
