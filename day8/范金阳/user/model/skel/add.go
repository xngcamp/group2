package skel

import "fmt"

// Register 定义新增操作
func (skel *Skel) Add() (err error) {
	c := skel.GetC()
	fmt.Printf("c:%v \n",c)
	defer c.Database.Session.Close()
	//arr :=make([]*Skel,0)
	//arr = append(arr,user)
	err = c.Insert(skel)
	if err != nil {
		return
	}

	return
}
