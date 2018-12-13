package feed

//Add 定义新增操作
func (feed *Feed) Add() (err error) {
	//连接数据库
	c := feed.GetC()
	//关闭数据库
	defer c.Database.Session.Close()

	//插入feed
	err = c.Insert(feed)
	if err != nil{
		return
	}
	return
}
