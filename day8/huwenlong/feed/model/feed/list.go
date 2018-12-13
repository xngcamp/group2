package feed

func (feed *Feed) List()(feeds []*Feed,err error)  {
	c := feed.GetC()
	defer c.Database.Session.Close()
	//返回参数为指针时 记得首先初始化避免返回空指针
	feeds = make([]*Feed,0)
	//获取全部feed
	err = c.Find(nil).All(&feeds)

	return
}