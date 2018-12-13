package skel

import (
	"testing"

	"camp/Execample/test"
)

// 测试/user/del
func TestDel(t *testing.T) {
	test.Setup()

	skel := GetSkel()
	id := skel.ID

	if err := Del(id); err != nil {
		t.Fatal(err)
	}

	skelNew, err := Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if skelNew != nil {
		t.Fatal("ret user not valid")
	}
}
