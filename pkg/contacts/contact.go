package contacts

// Contact provides the master structure for a contact
type Contact struct {
	ID         string `json:"id"`
	LastName   string `json:"lastname"`
	FirstName  string `json:"firstname"`
	BirthYear  string `json:"birthyear"`
	BirthMonth string `json:"birthmonth"`
	BirthDay   string `json:"birthday"`
	Email      string `json:"email"`
}
