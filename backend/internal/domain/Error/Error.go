package Error

import (
	"fmt"

	"admin/internal/utils/log"
)

type Error struct {
	code uint16
	msg  string
}

func (e Error) Wrap(err error) error {
	return fmt.Errorf("rpc error: %v, %s", e, err.Error())
}

func (e Error) LogWithTrace(err error) error {
	return fmt.Errorf("rpc error: %v, %s", e, log.WithTrace(err).Error())
}

var (
	OK                 = Error{code: 0, msg: ""}
	Canceled           = Error{code: 1, msg: ""}
	Unknown            = Error{code: 2, msg: ""}
	InvalidArgument    = Error{code: 3, msg: ""}
	DeadlineExceeded   = Error{code: 4, msg: ""}
	NotFound           = Error{code: 5, msg: ""}
	AlreadyExists      = Error{code: 6, msg: ""}
	PermissionDenied   = Error{code: 7, msg: ""}
	ResourceExhausted  = Error{code: 8, msg: ""}
	FailedPrecondition = Error{code: 9, msg: ""}
	Aborted            = Error{code: 10, msg: ""}
	OutOfRange         = Error{code: 11, msg: ""}
	Unimplemented      = Error{code: 12, msg: ""}
	Internal           = Error{code: 13, msg: ""}
	Unavailable        = Error{code: 14, msg: ""}
	DataLoss           = Error{code: 15, msg: ""}
	Unauthenticated    = Error{code: 16, msg: ""}
)
