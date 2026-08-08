package Test_Examples

import (
	"gopurs/output/gopurs_runtime"
	"time"
)

func SetTimeout(ms interface{}, f func(interface{}) interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		go func() {
			time.Sleep(time.Duration(ms.(int)) * time.Millisecond)
			f(nil)
		}()
		return nil
	}
}
