package builtin

import (
	"strings"

	. "github.com/apmckinlay/gsuneido/runtime"
)

func init() {
	SuFuncMethods = Methods{
		"Disasm": method0(func(this Value) Value {
			fn := this.(*SuFunc)
			buf := &strings.Builder{}
			Disasm(buf, fn)
			return SuStr(buf.String())
		}),
	}
}
