package contacts

// Contact provides the master structure for a contact
type Contact struct {
	ID         string `json:"id"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	BirthYear  int    `json:"birthyear"`
	BirthMonth int    `json:"birthmonth"`
	BirthDay   int    `json:"birthday"`
}
