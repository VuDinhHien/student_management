package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
}

type Student struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func GetAllStudents() ([]Student, error) {
	rows, err := db.Query("SELECT id, name, age, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func CreateStudent(student Student) error {
	_, err := db.Exec("INSERT INTO students (name, age, email) VALUES (?, ?, ?)", student.Name, student.Age, student.Email)
	return err
}

func GetStudentByID(id int) (Student, error) {
	var student Student
	err := db.QueryRow("SELECT id, name, age, email FROM students WHERE id = ?", id).Scan(&student.ID, &student.Name, &student.Age, &student.Email)
	return student, err
}

func UpdateStudent(student Student) error {
	_, err := db.Exec("UPDATE students SET name=?, age=?, email=? WHERE id=?", student.Name, student.Age, student.Email, student.ID)
	return err
}

func DeleteStudent(id int) error {
	_, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}
