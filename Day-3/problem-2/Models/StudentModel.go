package Models

type Student struct {
	ID           int           `json:"id"`
	FirstName    string        `json:"firstname"`
	LastName     string        `json:"lastname"`
	DOB          string        `json:"dob"`
	Address      string        `json:"address"`
	SubjectMarks []SubjectMark `gorm:"foreignKey:ID"`
}

type SubjectMark struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
}

func (b *Student) TableName() string {
	return "Student Table"
}

func (b *SubjectMark) TableName() string {
	return "Subject Marks Table"
}
