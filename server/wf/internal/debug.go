package internal

import (
	"fmt"
	"runtime"
)

const (
	roundMark = "----------------"
)

func callerName() string {
	funName := "<Unknown>"
	pc, _, _, ok := runtime.Caller(3)
	if ok {
		funName = runtime.FuncForPC(pc).Name()
	}
	return funName
}

func DebugMsg(msg any) {
	DebugTitleMsg(callerName(), msg)
}

func DebugTitleMsg(title, msg any) {
	if GlobalConfig.Debug {
		fmt.Printf("%s%s%s\n", roundMark, title, roundMark)
		fmt.Printf("%+v\n", msg)
	}
}
