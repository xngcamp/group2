package session

func (s *Session) Add() (err error) {
	c := s.GetC()
	defer c.Database.Session.Close()
	err = c.Insert(s)
	if err != nil {
		return
	}
	return
}

