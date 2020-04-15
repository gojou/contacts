package contacts

// Service defines the methods supported by the Contacts package
type Service interface {
	Create(Contact) string
	RetrieveList() []Contact
}

// Repo defines the calls to the data repository
type Repo interface {
	Create(Contact) string
	RetrieveList() []Contact
}

type service struct {
	r Repo
}

// NewService returns the Contact service
func NewService(r Repo) Service {
	return &service{r}
}

func (s *service) Create(c Contact) string {
	return s.r.Create(c)
}

func (s *service) RetrieveList() []Contact {
	return s.r.RetrieveList()
}
