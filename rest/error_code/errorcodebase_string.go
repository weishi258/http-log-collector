// Code generated by "stringer -type=ErrorCodeBase"; DO NOT EDIT.

package error_code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InternalError-0]
	_ = x[InvalidParameter-1]
	_ = x[BadRequest-2]
}

const _ErrorCodeBase_name = "InternalErrorInvalidParameterBadRequest"

var _ErrorCodeBase_index = [...]uint8{0, 13, 29, 39}

func (i ErrorCodeBase) String() string {
	if i < 0 || i >= ErrorCodeBase(len(_ErrorCodeBase_index)-1) {
		return "ErrorCodeBase(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCodeBase_name[_ErrorCodeBase_index[i]:_ErrorCodeBase_index[i+1]]
}
