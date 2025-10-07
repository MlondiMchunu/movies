package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// define an error that UnmarshalJSON() method can return
// if JSON string cant be converted or parsed
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// UnmarshalJSON() method satisfies the
// json.Unmarshaler interface
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	//remove surrounding double-quotes from string.

	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))

	// If  unquoting isn't possible, then return
	// ErrInvalidRuntimeFormat error.
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, "")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
}
