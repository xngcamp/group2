// Code generated by wsp, DO NOT EDIT.

package main

import "net/http"
import "time"
import "camp/feed/controller/feed"
import "camp/feed/filter"

func init() {
	http.HandleFunc("/feed/add", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/add"}); !ok {
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
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/del"}); !ok {
				return
			}
		}()
		c.Del(w, r)
	})

	http.HandleFunc("/feed/get", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/get"}); !ok {
				return
			}
		}()
		c.Get(w, r)
	})

	http.HandleFunc("/feed/list", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
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
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/update"}); !ok {
				return
			}
		}()
		c.Update(w, r)
	})

}