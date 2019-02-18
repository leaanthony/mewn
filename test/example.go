package test

import (
	"fmt"

	"github.com/leaanthony/mewn"
)

func mytest() {
	myasset1 := mewn.String("./assets/hello.txt")
	myasset2 := mewn.String("./assets/hello.txt")
	fmt.Printf("myasset1 = '%s'\n", myasset1)
	fmt.Printf("myasset2 = '%s'\n", myasset2)
}
