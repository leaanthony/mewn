package mewn

import (
	"fmt"
)

func main() {
	mygroup := Group("./assets")
	myasset := mygroup.String("hello.txt")
	fmt.Printf("myasset = '%s'\n", myasset)
}
