package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)      //darwin
	fmt.Println(runtime.GOARCH)    //arm64
	fmt.Println(runtime.NumCPU())  //8
	fmt.Println(runtime.Version()) // go1.22.3
}
