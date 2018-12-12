package user

func (u *User) Update() (err error) {
	c := u.GetC()
	err = c.UpdateId(u.Id, u)
	if err != nil {
		return
	}
	return
}
