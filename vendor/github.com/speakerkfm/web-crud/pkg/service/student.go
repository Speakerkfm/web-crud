package service

import (
	"github.com/speakerkfm/web-crud/pkg/store"
	"github.com/speakerkfm/web-crud/rpc/psu"
	"github.com/twitchtv/twirp"
)

type StudentService struct {
	st store.StoreInterface
}

func NewStudentService(st store.StoreInterface) *StudentService {
	return &StudentService{st: st}
}

func (ss *StudentService) GetStudentsList() ([]*psu.Student, error) {
	students, err := ss.st.GetStudents()
	if err != nil {
		return nil, err
	}

	studentsList := make([]*psu.Student, len(students))
	for idx, student := range students {
		studentsList[idx] = &psu.Student{
			Id:      student.ID,
			Name:    student.Name,
			Surname: student.Surname,
		}
	}

	return studentsList, nil
}

func (ss *StudentService) CreateStudent(person *psu.Person) (*psu.Student, error) {
	student, err := ss.st.CreateStudent(person.Name, person.Surname)
	if err != nil {
		return nil, err
	}

	return &psu.Student{
		Id:      student.ID,
		Name:    student.Name,
		Surname: student.Surname,
	}, nil
}

func (ss *StudentService) DeleteStudent(studentID *psu.StudentID) error {
	if _, err := ss.st.GetStudentByID(studentID.Id); err != nil {
		return twirp.NotFoundError(err.Error())
	}

	return ss.st.DeleteStudent(studentID.Id)
}

func (ss *StudentService) UpdateStudent(student *psu.Student) (*psu.Student, error) {
	oldStudent, err := ss.st.GetStudentByID(student.Id)
	if err != nil {
		return nil, twirp.NotFoundError(err.Error())
	}

	if student.Name == "" {
		student.Name = oldStudent.Name
	}
	if student.Surname == "" {
		student.Surname = oldStudent.Surname
	}

	updatedStudent, err := ss.st.UpdateStudent(student.Id, student.Name, student.Surname)
	if err != nil {
		return nil, err
	}

	return &psu.Student{
		Id:      updatedStudent.ID,
		Name:    updatedStudent.Name,
		Surname: updatedStudent.Surname,
	}, nil
}
