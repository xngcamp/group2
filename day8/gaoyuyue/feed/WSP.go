// Code generated by wsp, DO NOT EDIT.

package main

import "net/http"
import "time"
import "github.com/xngcamp/group2/day8/gaoyuyue/feed/controller/feed"
import "github.com/xngcamp/group2/day8/gaoyuyue/feed/controller/user"
import "github.com/xngcamp/group2/day8/gaoyuyue/feed/filter"

func init() {
	http.HandleFunc("/feed/add", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Cors(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/add"}); !ok {
				return
			}
		}()
		c.Add(w, r)
	})

	http.HandleFunc("/feed/del", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Cors(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/del"}); !ok {
				return
			}
		}()
		c.Del(w, r)
	})

	http.HandleFunc("/feed/list", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Cors(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/list"}); !ok {
				return
			}
		}()
		c.List(w, r)
	})

	http.HandleFunc("/feed/update", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Cors(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/update"}); !ok {
				return
			}
		}()
		c.Update(w, r)
	})

	http.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/login"}); !ok {
				return
			}
		}()
		c.Login(w, r)
	})

	http.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/register"}); !ok {
				return
			}
		}()
		c.Register(w, r)
	})

	http.HandleFunc("/user/sub_user", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/sub_user"}); !ok {
				return
			}
		}()
		c.SubUser(w, r)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/update"}); !ok {
				return
			}
		}()
		if ok := filter.Auth(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/update"}); !ok {
			return
		}
		c.Update(w, r)
	})

}