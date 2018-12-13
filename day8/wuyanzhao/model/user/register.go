package user

// Register 定义用户注册操作
func (userModel *UserModel) Register() (err error) {
	c := userModel.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(userModel)
	if err != nil {
		return
	}

	return
}
