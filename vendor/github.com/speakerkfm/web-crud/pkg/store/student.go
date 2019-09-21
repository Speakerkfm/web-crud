package store

type Student struct {
	ID      int32 `gorm:"primary_key"`
	Name    string
	Surname string
}

func (st *Store) GetStudentByID(studentID int32) (*Student, error) {
	student := &Student{ID: studentID}

	err := st.gorm.First(student).Error

	return student, err
}

func (st *Store) GetStudents() ([]Student, error) {
	var studentsList []Student

	err := st.gorm.Table("student").Scan(&studentsList).Error

	return studentsList, err
}

func (st *Store) CreateStudent(name, surname string) (*Student, error) {
	student := Student{
		Name:    name,
		Surname: surname,
	}

	err := st.gorm.Create(&student).Error

	return &student, err
}

func (st *Store) DeleteStudent(studentID int32) error {
	return st.gorm.Delete(&Student{ID: studentID}).Error
}

func (st *Store) UpdateStudent(id int32, name, surname string) (*Student, error) {
	student := Student{
		ID:      id,
		Name:    name,
		Surname: surname,
	}

	err := st.gorm.Save(&student).Error

	return &student, err
}
