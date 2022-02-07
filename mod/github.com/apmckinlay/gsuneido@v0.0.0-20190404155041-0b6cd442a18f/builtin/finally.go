package builtin

import (
	. "github.com/apmckinlay/gsuneido/runtime"
)

var _ = builtinRaw("Finally(main_block, final_block)",
	func(t *Thread, as *ArgSpec, args ...Value) Value {
		args = t.Args(&paramSpecFinally, as)
		defer func() {
			e := recover()
			func() {
				defer func() {
					if e != nil {
						recover() // if main block panics, ignore finally panic
					}
				}()
				t.CallWithArgs(args[1])
			}()
			if e != nil {
				panic(e)
			}
		}()
		return t.CallWithArgs(args[0])
	})

var paramSpecFinally = params(`(main_block, final_block)`)
