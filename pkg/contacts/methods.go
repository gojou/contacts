package contacts

func (s *service) Create(c Contact) string {
	return s.r.Create(c)
}

func (s *service) RetrieveList() []Contact {
	return s.r.RetrieveList()
}
