package feed

func (f *Feed) Update() error {
	c := f.GetC()
	defer c.Database.Session.Close()
	return c.UpdateId(f.Id, f)
}
