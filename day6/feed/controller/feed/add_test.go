package feed

import (
	"testing"

	"camp/feed/test"
)

// 测试/feed/add
func TestAdd(t *testing.T) {
	test.Setup()

	feed := GetFeed()
	if feed == nil {
		t.Fatal("test fail")
	}
}
