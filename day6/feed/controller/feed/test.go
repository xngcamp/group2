// 此文件封装仅用于单元测试公共函数

package feed

import (
	"encoding/json"
	"fmt"
	"strconv"

	"camp/feed/api"
	"camp/feed/test"
	"camp/lib"
)

// Add 封装controller.Add操作
func Add(id int) (err error) {
	c := &Feed{}
	p := &AddReq{
		ID: strconv.Itoa(id),
	}
	body, err := lib.TestPost(c.Add, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data AddResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	test.Sleep()
	return
}

// Del 封装controller.Del操作
func Del(id int) (err error) {
	c := &Feed{}
	p := &DelReq{
		ID: strconv.Itoa(id),
	}
	body, err := lib.TestPost(c.Del, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data AddResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	test.Sleep()
	return
}

// Get 封装controller.Get操作
func Get(id int) (feedApi *api.Feed, err error) {
	c := &Feed{}
	p := &GetReq{
		ID: strconv.Itoa(id),
	}
	body, err := lib.TestPost(c.Get, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data GetResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	feedApi = resp.Data.Feed
	return
}

// GetSkel 返回一个全新的skel对象
func GetFeed() (feedApi *api.Feed) {
	id := test.GetID()
	for idTemp := id; ; idTemp++ {
		if feedTemp, err := Get(idTemp); err != nil {
			panic(err)
		} else if feedTemp != nil {
			if err := Del(idTemp); err != nil {
				panic(err)
			}
		} else {
			break
		}
	}

	if err := Add(id); err != nil {
		panic(err)
	}

	feedApi, err := Get(id)
	if err != nil {
		panic(err)
	} else if feedApi == nil {
		panic("get feed empty")
	}

	return
}
