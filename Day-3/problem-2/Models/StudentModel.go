package Models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
}

//type SubjectMarks struct {
//	ID      int    `json:"id"`
//	Subject string `json:"subject"`
//	Marks   int    `json:"marks"`
//}

func (b *Student) TableName() string {
	return "Student Table"
}

//func (b *SubjectMarks) TableName() string {
//	return "Subject Marks Table"
//}
