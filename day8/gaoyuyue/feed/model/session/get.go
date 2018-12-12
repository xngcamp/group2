package session

func (s *Session) Get() (err error) {
	c := s.GetC()
	defer c.Database.Session.Close()
	err = c.FindId(s.Id).One(s)
	if err != nil {
		return
	}
	return
}
