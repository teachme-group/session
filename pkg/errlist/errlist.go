package errlist

import "github.com/Markuysa/pkg/errs"

var (
	ErrUnauthorized = errs.New("unauthorized", errs.Unauthenticated)
)
