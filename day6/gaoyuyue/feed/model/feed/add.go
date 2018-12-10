package feed

func (f *Feed) Add() error {
	c := f.GetC()
	defer c.Database.Session.Close()
	return c.Insert(f)
}
