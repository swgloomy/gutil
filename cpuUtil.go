package gutil

import "runtime"

/**
设置CPU核心使用数
创建人:邵炜
创建时间:2021-03-18 11:13:25
*/
func SetCPUUseNumber(number int) {
	numberCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(If(number <= numberCPU && number > 0, number, numberCPU).(int))
}
