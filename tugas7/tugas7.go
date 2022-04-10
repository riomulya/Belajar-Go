package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	x := 1
	y := "string"
	runtime.GOMAXPROCS(2)
	go bacaTipeInt(x)
	bacaTipeStr(y)

	var print string
	fmt.Scanln(&print)
}

func bacaTipeInt(bacaint int) {
	fmt.Println("Nilainya Adalah :", reflect.ValueOf(bacaint))
	fmt.Println("Tipenyanya Adalah :", reflect.ValueOf(bacaint).Type())
}
func bacaTipeStr(bacaStr string) {
	fmt.Println("Nilainya Adalah :", reflect.ValueOf(bacaStr))
	fmt.Println("Tipenyanya Adalah :", reflect.ValueOf(bacaStr).Type())
}
