package skel

import (
	"testing"

	"camp/Execample/test"
)

// 测试/user/get
func TestGet(t *testing.T) {
	test.Setup()

	skel := GetSkel()
	if skel == nil {
		t.Fatal("test fail")
	}
}
