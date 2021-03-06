// Code generated by "stringer -type=Type"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Boolean-0]
	_ = x[Number-1]
	_ = x[String-2]
	_ = x[Date-3]
	_ = x[Object-4]
	_ = x[Record-5]
	_ = x[Function-6]
	_ = x[Block-7]
	_ = x[BuiltinFunction-8]
	_ = x[Class-9]
	_ = x[Method-10]
	_ = x[Except-11]
	_ = x[Instance-12]
	_ = x[Iterator-13]
}

const _Type_name = "BooleanNumberStringDateObjectRecordFunctionBlockBuiltinFunctionClassMethodExceptInstanceIterator"

var _Type_index = [...]uint8{0, 7, 13, 19, 23, 29, 35, 43, 48, 63, 68, 74, 80, 88, 96}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
