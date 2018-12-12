package user

func (u *User) Add() (err error) {
	c := u.GetC()
	defer c.Database.Session.Close()
	if err = c.Insert(u); err != nil {
		return
	}
	return
}
