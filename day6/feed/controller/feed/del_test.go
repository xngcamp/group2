package feed

import (
	"strconv"
	"testing"

	"camp/feed/test"
)

// 测试/feed/del
func TestDel(t *testing.T) {
	test.Setup()

	feed := GetFeed()
	id, _ := strconv.Atoi(feed.ID)

	if err := Del(id); err != nil {
		t.Fatal(err)
	}

	feedNew, err := Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if feedNew != nil {
		t.Fatal("ret feed not valid")
	}
}
