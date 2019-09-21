package psuserver

import (
	"context"
	"github.com/speakerkfm/web-crud/rpc/psu"
	"github.com/twitchtv/twirp"
)

func (s *Server) CreateStudent(ctx context.Context, person *psu.Person) (*psu.Student, error) {
	return s.StudentService.CreateStudent(person)
}

func (s *Server) GetStudentsList(ctx context.Context, empty *psu.Empty) (*psu.StudentsList, error) {
	studentsList, err := s.StudentService.GetStudentsList()
	if err != nil {
		return nil, err
	}

	return &psu.StudentsList{Students: studentsList}, nil
}

func (s *Server) UpdateStudent(ctx context.Context, student *psu.Student) (*psu.Student, error) {
	if student.Id <= 0 {
		return nil, twirp.InvalidArgumentError("id", "must be unsigned int")
	}
	return s.StudentService.UpdateStudent(student)
}

func (s *Server) DeleteStudent(ctx context.Context, studentID *psu.StudentID) (*psu.Empty, error) {
	if studentID.Id <= 0 {
		return nil, twirp.InvalidArgumentError("id", "must be unsigned int")
	}
	return &psu.Empty{}, s.StudentService.DeleteStudent(studentID)
}
