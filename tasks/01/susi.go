package main

type Susi struct {
}

type Course struct {
	courseIdentifier string
	courseName string
	minimumAcademicYear int
	mastersOnly bool
	availablePlaces int
}

func (c *Course) String() string {
	return c.courseIdentifier + " " + c.courseName
} 

type Student struct {
	FacultyNumber int
	FirstName string
	LastName string
	AcademicYear int
	Master bool
}

func (s *Student) String() string {
	return s.FacultyNumber + " " + s.FirstName + " " + s.LastName
}

type SusiError struct {
	student Student
	course Course = nil	
}

type Enrollment struct {
	course Course
	student Student
}

func NewSusi() *Susi {
	return new 
}

func (s *Susi) AddStudent(request []byte) error
func (s *Susi) FindStudent(facultyNumber int) (*Student, error)
func (s *Susi) AddCourse(request []byte) error
func (s *Susi) FindCourse(courseIdentifier string) (*Course, error)
func (s *Susi) Enroll(request []byte) error
func (s *Susi) FindEnrollment(facultyNumber int, courseIdentifier string)
