package etc

import (
	"bytes"
	"fmt"
)

// given a run of errors, returns the first non-nil error.
func FirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// given a run of errors, merge all of the non-nil error values into a single
// error value.
func MergeErrors(errs ...error) error {
	var l *ErrorList

	for _, err := range errs {
		if err == nil {
			continue
		}

		if l == nil {
			l = &ErrorList{vals: make([]error, 0, len(errs))}
		}

		l.vals = append(l.vals, err)
	}

	if l == nil {
		return nil
	}
	switch len(l.vals) {
	case 0:
		// can we even get here?
		return nil
	case 1:
		return l.vals[0]
	default:
		return l
	}
}

type ErrorList struct{ vals []error }

func (e *ErrorList) Error() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%d errors:", len(e.vals))
	for _, err := range e.vals {
		fmt.Fprintf(&b, " [%v]", err)
	}

	return b.String()
}

func (e *ErrorList) Errors() []error { return e.vals }

func TossValue(v interface{}, err error) error { return err }
