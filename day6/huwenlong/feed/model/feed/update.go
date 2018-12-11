package feed

//Add 定义新增操作
func (feed *Feed) Update() (err error) {
	c := feed.GetC()

	defer c.Database.Session.Close()

	err = c.UpdateId(feed.Id,feed)
	if err != nil{
		return
	}
	return
}
