package errors

import (
	"fmt"
	"net/http"
)

type Err = uint16

const (
	ErrorInvalidQuery uint16 = iota
	ErrorBadLoginOrPassword
	ErrorUnauthorized
	ErrorCannotAccessRcon
)

type McapError struct {
	err         uint16
	stringified string
}

func New(e Err) McapError {
	var stringified string

	if e == ErrorInvalidQuery {
		stringified = "invalid query"
	}
	if e == ErrorBadLoginOrPassword {
		stringified = "bad login or password"
	}
	return McapError{
		stringified: stringified,
		err:         e,
	}
}

func (e *McapError) Stringify() string {
	return e.stringified
}

func (e *McapError) IsEqualRaw(err Err) bool {
	return e.err == err
}

func (e *McapError) IsEqual(err McapError) bool {
	return err.err == e.err
}

func HttpError(w http.ResponseWriter, err Err, code int) {
	w.WriteHeader(code)

	w.Write([]byte(
		fmt.Sprintf(
			`{"err":"%v"}`,
			New(err).err, //.stringified,
		),
	))
}
