package feed

func (f *Feed) List() (feeds []*Feed, err error) {
	c := f.GetC()
	defer c.Database.Session.Close()
	feeds = make([]*Feed, 0)
	err = c.Find(nil).All(&feeds)
	return
}
