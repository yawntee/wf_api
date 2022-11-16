package internal

import "runtime"

func FuncName() string {
	funName := "<Unknown>"
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		funName = runtime.FuncForPC(pc).Name()
	}
	return funName
}
