// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package opcodes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Nop-0]
	_ = x[Pop-1]
	_ = x[Dup-2]
	_ = x[Dup2-3]
	_ = x[Dupx2-4]
	_ = x[Int-5]
	_ = x[Value-6]
	_ = x[True-7]
	_ = x[False-8]
	_ = x[Zero-9]
	_ = x[One-10]
	_ = x[MaxInt-11]
	_ = x[EmptyStr-12]
	_ = x[Load-13]
	_ = x[Store-14]
	_ = x[Dyload-15]
	_ = x[Global-16]
	_ = x[Get-17]
	_ = x[Put-18]
	_ = x[RangeTo-19]
	_ = x[RangeLen-20]
	_ = x[This-21]
	_ = x[Is-22]
	_ = x[Isnt-23]
	_ = x[Match-24]
	_ = x[MatchNot-25]
	_ = x[Lt-26]
	_ = x[Lte-27]
	_ = x[Gt-28]
	_ = x[Gte-29]
	_ = x[Add-30]
	_ = x[Sub-31]
	_ = x[Cat-32]
	_ = x[Mul-33]
	_ = x[Div-34]
	_ = x[Mod-35]
	_ = x[LeftShift-36]
	_ = x[RightShift-37]
	_ = x[BitOr-38]
	_ = x[BitAnd-39]
	_ = x[BitXor-40]
	_ = x[BitNot-41]
	_ = x[Not-42]
	_ = x[UnaryPlus-43]
	_ = x[UnaryMinus-44]
	_ = x[Or-45]
	_ = x[And-46]
	_ = x[Bool-47]
	_ = x[QMark-48]
	_ = x[In-49]
	_ = x[Jump-50]
	_ = x[JumpTrue-51]
	_ = x[JumpFalse-52]
	_ = x[JumpIs-53]
	_ = x[JumpIsnt-54]
	_ = x[Iter-55]
	_ = x[ForIn-56]
	_ = x[Throw-57]
	_ = x[Try-58]
	_ = x[Catch-59]
	_ = x[CallFunc-60]
	_ = x[CallMeth-61]
	_ = x[Super-62]
	_ = x[Return-63]
	_ = x[ReturnNil-64]
	_ = x[Block-65]
	_ = x[BlockBreak-66]
	_ = x[BlockContinue-67]
	_ = x[BlockReturn-68]
	_ = x[BlockReturnNil-69]
}

const _Opcode_name = "NopPopDupDup2Dupx2IntValueTrueFalseZeroOneMaxIntEmptyStrLoadStoreDyloadGlobalGetPutRangeToRangeLenThisIsIsntMatchMatchNotLtLteGtGteAddSubCatMulDivModLeftShiftRightShiftBitOrBitAndBitXorBitNotNotUnaryPlusUnaryMinusOrAndBoolQMarkInJumpJumpTrueJumpFalseJumpIsJumpIsntIterForInThrowTryCatchCallFuncCallMethSuperReturnReturnNilBlockBlockBreakBlockContinueBlockReturnBlockReturnNil"

var _Opcode_index = [...]uint16{0, 3, 6, 9, 13, 18, 21, 26, 30, 35, 39, 42, 48, 56, 60, 65, 71, 77, 80, 83, 90, 98, 102, 104, 108, 113, 121, 123, 126, 128, 131, 134, 137, 140, 143, 146, 149, 158, 168, 173, 179, 185, 191, 194, 203, 213, 215, 218, 222, 227, 229, 233, 241, 250, 256, 264, 268, 273, 278, 281, 286, 294, 302, 307, 313, 322, 327, 337, 350, 361, 375}

func (i Opcode) String() string {
	if i >= Opcode(len(_Opcode_index)-1) {
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}