package token

func (t *Token) Set() (err error) {
	conn := t.GetConn()
	if _, err = conn.Do("SET", t.Id.Hex(), t.UserId.Hex()); err != nil {
		return
	}
	return
}