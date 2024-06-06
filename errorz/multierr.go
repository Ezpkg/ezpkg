package errorz

import (
	"fmt"
	"reflect"

	"github.com/ezpkg/stacktracez"
)

var _ Errors = (*zErrors)(nil)
var _ stacktracez.StackTracerZ = (*zErrors)(nil)

type Errors interface {
	error
	Errors() []error
}

func Append(pErr *error, errs ...error) {
	appendErrs(Option{}, pErr, errs...)
}

func appendErrs(opt Option, pErr *error, errs ...error) {
	switch err0 := (*pErr).(type) {
	case nil:
		var zErrs zErrors
		zErrs.Append(errs...)
		*pErr = zErrs.process(opt)

	case *zErrors:
		err0.Append(errs...)
		*pErr = err0.process(opt)

	default:
		var zErrs zErrors
		zErrs.errors = make([]error, len(errs))
		zErrs.Append(errs...)
		*pErr = zErrs.process(opt)
	}
}

func Appendf(pErr *error, err error, msg string, args ...any) {
	if err == nil {
		return
	}
	if msg != "" {
		err = Wrapf(err, msg, args...)
	}
	appendErrs(Option{}, pErr, err)
}

func Validatef(pErr *error, condition bool, msg string, args ...any) {
	if !condition {
		appendErrs(Option{}, pErr, NoStack().Errorf(msg, args...))
	}
}

func ValidateX[T any](pErr *error, value T, condition bool, msg string, args ...any) (out T) {
	if condition {
		return value
	} else {
		appendErrs(Option{}, pErr, NoStack().Errorf(msg, args...))
		return out
	}
}

func GetErrors(err error) []error {
	switch err := err.(type) {
	case interface{ Errors() []error }:
		if err != nil {
			return err.Errors()
		}
	case interface{ WrappedErrors() []error }:
		if err != nil {
			return err.WrappedErrors()
		}
	}
	return nil
}

type zErrors struct {
	errors []error // len(errors) > 0
	stack  *stacktracez.Frames
}

func (es *zErrors) Error() string {
	switch {
	case es == nil:
		return "<nil>"
	case len(es.errors) == 0:
		return "<empty>"
	case len(es.errors) == 1:
		return fmt.Sprintf("(1 error) %v", es.errors[0])
	default:
		return fmt.Sprintf("%v", es)
	}
}

func (es *zErrors) Format(s fmt.State, verb rune) {
	if es == nil {
		writeString(s, "<nil>")
		return
	}
	isPlus := s.Flag('+') || s.Flag('#')
	switch verb {
	case 's', 'v':
		switch {
		case len(es.errors) == 0:
			writeString(s, "<empty>")
			return
		case len(es.errors) == 1:
			if isPlus {
				fprintf(s, "1 error occurred:\n\t* %v\n", es.errors[0])
			} else {
				fprintf(s, "(1 error) %v", es.errors[0])
			}
			if isPlus && es.stack != nil {
				es.stack.Format(s, verb)
			}
		default:
			if isPlus {
				fprintf(s, "%d errors occurred:\n", len(es.errors))
			} else {
				fprintf(s, "(%d errors) ", len(es.errors))
			}
			for i, err := range es.errors {
				if isPlus {
					fprintf(s, "\t* %v\n", err)
					continue
				}
				if i > 0 {
					fprintf(s, " ; ")
				}
				fprintf(s, "%v", err)
			}
			if isPlus && es.stack != nil {
				es.stack.Format(s, verb)
			}
		}
	case 'd':
		fprintf(s, "%d", len(es.errors))
	}
}

func (es *zErrors) Errors() []error {
	if es == nil {
		return nil
	}
	return es.errors[:len(es.errors)]
}

func (es *zErrors) StackTraceZ() *stacktracez.Frames {
	if es == nil {
		return nil
	}
	return es.stack
}

func (es *zErrors) Append(errs ...error) {
	for _, err := range errs {
		switch err := err.(type) {
		case nil:
			// continue
		case *zErrors:
			if err != nil {
				es.errors = append(es.errors, err.errors...)
			}
		case interface{ Errors() []error }:
			// uber-go/multierr, tailscale.com/util/multierr
			if err != nil {
				es.errors = append(es.errors, err.Errors()...)
			}
		case interface{ WrappedErrors() []error }:
			// hashicorp/go-multierror
			if err != nil {
				es.errors = append(es.errors, err.WrappedErrors()...)
			}
		default:
			if !reflect.ValueOf(err).IsNil() {
				es.errors = append(es.errors, err)
			}
		}
	}
}

func (es *zErrors) Appendf(err error, msgArgs ...any) {
	err = formatMsg(err, msgArgs)
	if err != nil {
		es.errors = append(es.errors, err)
	}
}

func (es *zErrors) Validatef(condition bool, msg string, args ...any) {
	if !condition {
		err := NoStack().Errorf(msg, args...)
		es.errors = append(es.errors, err)
	}
}

func (es *zErrors) process(opt Option) error {
	if len(es.errors) == 0 {
		return nil
	}
	if es.stack != nil || opt.NoStack {
		return es
	}
	es.stack = stacktracez.StackTraceSkip(opt.CallersSkip + 3)
	return es
}

func formatMsg(err error, msgArgs []any) error {
	if len(msgArgs) == 0 {
		return err
	}
	if format, ok := msgArgs[0].(string); ok {
		return NoStack().Wrapf(err, format, msgArgs[1:]...)
	}
	return NoStack().Wrap(err, fmt.Sprintln(msgArgs...))
}