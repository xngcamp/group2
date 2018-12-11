package feed

func (feed *Feed) List()(feeds []*Feed,err error)  {
	c := feed.GetC()
	defer c.Database.Session.Close()

	feeds = make([]*Feed,0)
	//获取全部feed
	err = c.Find(nil).All(&feeds)

	return
}