package test

import (
	"fmt"

	"github.com/leaanthony/mewn"
)

func mytest() {
	myasset := mewn.String("./assets/hello.txt")
	fmt.Printf("myasset = '%s'\n", myasset)
}
