package Models

type Student struct {
	ID           int           //`json:"id"`
	FirstName    string        //`json:"firstname"`
	LastName     string        //`json:"lastname"`
	DOB          string        //`json:"dob"`
	Address      string        //`json:"address"`
	SubjectMarks []SubjectMark //`json:"subject_marks"`
}

type SubjectMark struct {
	ID        int    //`json:"id"`
	Subject   string //`json:"subject"`
	Marks     int    //`json:"marks"`
	StudentID int    //`json:"student_id"`
}

func (b *Student) TableName() string {
	return "student"
}

func (b *SubjectMark) TableName() string {
	return "subjectmark"
}
