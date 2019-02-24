package test

import (
	"fmt"

	"github.com/leaanthony/mewn"
)

func mygrouptest() {
	myGroup := mewn.Group("./assets")
	myasset1 := myGroup.String("hello.txt")
	fmt.Printf("myasset1 = '%s'\n", myasset1)
}
