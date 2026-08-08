package Test_Bench

import (
	"runtime"
)

func Gc(_ interface{}) interface{} {
	runtime.GC()
	return nil
}
