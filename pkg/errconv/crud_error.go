package errconv

import "errors"

var (
	ErrNotFound = errors.New("entity not found")
)

// type ErrNotFound struct {
// 	Msg string
// }

// func (err *ErrNotFound) Error() string {
// 	return err.Msg
// }

// func (err *ErrNotFound) New() error {
// 	return errors.New("entity not found")
// }
