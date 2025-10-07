package data

import (
	"errors"
	"fmt"
	"strconv"
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
