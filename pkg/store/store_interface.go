// Code generated by ifacemaker. DO NOT EDIT.

package store

type StoreInterface interface {
	GetStudentByID(studentID int32) (*Student, error)
	GetStudents() ([]Student, error)
	CreateStudent(name, surname string) (*Student, error)
	DeleteStudent(studentID int32) error
	UpdateStudent(id int32, name, surname string) (*Student, error)
}
