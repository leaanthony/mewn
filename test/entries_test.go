package test

import (
	"testing"

	"github.com/leaanthony/mewn"
)

func TestSingleEntryList(t *testing.T) {
	mewn.AddAsset("./assets/hello.txt", "test")
	list := mewn.Entries()
	if list[0] != "./assets/hello.txt" {
		t.Fail()
	}
}

func TestMultipleEntryList(t *testing.T) {
	mewn.Reset()
	mewn.AddAsset("./assets/hello.txt", "test")
	mewn.AddAsset("./assets/hello2.txt", "test")
	list := mewn.Entries()
	for _, entry := range list {
		if entry != "./assets/hello.txt" && entry != "./assets/hello2.txt" {
			t.Fail()
		}
	}
}
