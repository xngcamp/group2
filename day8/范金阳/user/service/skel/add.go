package skel

import (
	"camp/Execample/api"
	"camp/Execample/model"
	"fmt"
)

// Register 定义新增操作
func (skel *Skel) Add(skelApi *api.Skel) (err error) {
	skelModel := model.NewSkel()
	skelModel.ID = skelApi.ID
	fmt.Printf("service:%v", skelModel)
	if err = skelModel.Add(); err != nil {
		return
	}

	return
}
