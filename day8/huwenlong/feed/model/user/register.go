package user

func (user *User) Register() (err error) {
	c := user.GetC()

	defer c.Database.Session.Close()
	err = c.Insert(user)
	if err != nil {
		return
	}
	return
}
