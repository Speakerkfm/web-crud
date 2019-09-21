package main

import (
	"context"
	"fmt"
	"github.com/speakerkfm/web-crud/rpc/psu"
	"net/http"
	"os"
)

func main() {
	client := psu.NewPsuProtobufClient("http://localhost:8080", &http.Client{})

	//CREATE
	student, err := client.CreateStudent(context.Background(), &psu.Person{Name: "Alexandr", Surname: "Usanin"})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Welcome new student %v %s %s \n", student.Id, student.Name, student.Surname)

	//READ
	studentsList, err := client.GetStudentsList(context.Background(), &psu.Empty{})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	for _, studentInList := range studentsList.Students {
		fmt.Println(studentInList)
	}

	//DELETE
	_, err = client.DeleteStudent(context.Background(), &psu.StudentID{Id: student.Id})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Student %v has been deleted \n", student.Id)

	studentsList, err = client.GetStudentsList(context.Background(), &psu.Empty{})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	for _, studentInList := range studentsList.Students {
		fmt.Println(studentInList)
	}

	//UPDATE
	updatedStudent, err := client.UpdateStudent(context.Background(), &psu.Student{Id: 15, Name: "kek"})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Updated student %v %s %s \n", updatedStudent.Id, updatedStudent.Name, updatedStudent.Surname)
}
