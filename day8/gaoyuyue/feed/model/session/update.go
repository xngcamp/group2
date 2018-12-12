package session

func (s *Session) Update() (err error) {
	c := s.GetC()
	defer c.Database.Session.Close()
	err = c.UpdateId(s.Id, c)
	if err != nil {
		return
	}
	return
}
