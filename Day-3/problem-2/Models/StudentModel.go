package Models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Subject   string `json:"subject"`
	Marks     string `json:"marks"`
}

func (b *Student) TableName() string {
	return "student"
}
