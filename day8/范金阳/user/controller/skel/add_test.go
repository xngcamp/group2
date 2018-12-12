package skel

import (
	"testing"

	"camp/Execample/test"
)

// 测试/user/add
func TestAdd(t *testing.T) {
	test.Setup()

	skel := GetSkel()
	if skel == nil {
		t.Fatal("test fail")
	}
}
