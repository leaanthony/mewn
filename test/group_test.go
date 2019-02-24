package test

import (
	"fmt"
	"testing"

	"github.com/leaanthony/mewn/lib"
)

func TestGroup(t *testing.T) {
	mewnFiles := lib.GetMewnFiles("group.go")
	if len(mewnFiles) != 1 {
		t.Fail()
	}

	referencedAssets, err := lib.GetReferencedAssets(mewnFiles)
	if err != nil {
		t.Fail()
	}

	for _, ass := range referencedAssets {
		fmt.Printf("%+v\n", ass)
	}

	// if len(referencedAssets) != 1 {
	// 	t.Fail()
	// }

	// theAsset := referencedAssets[0]
	// packedFileString := lib.GeneratePackFileString(theAsset)
	// fixture, err := ioutil.ReadFile("./fixtures/example.go.txt")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// // Fix line endings for fixture
	// if runtime.GOOS == "windows" {
	// 	fixture = bytes.Replace(fixture, []byte{13, 10}, []byte{10}, -1)
	// }

	// if string(fixture) != packedFileString {
	// 	fmt.Println(string(fixture))
	// 	fmt.Println(packedFileString)
	// 	t.Fail()
	// }
}
