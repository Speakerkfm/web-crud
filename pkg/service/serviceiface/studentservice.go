// Code generated by ifacemaker. DO NOT EDIT.

package serviceiface

import (
	"github.com/speakerkfm/web-crud/rpc/psu"
)

type StudentService interface {
	GetStudentsList() ([]*psu.Student, error)
	CreateStudent(person *psu.Person) (*psu.Student, error)
	DeleteStudent(studentID *psu.StudentID) error
	UpdateStudent(student *psu.Student) (*psu.Student, error)
}
