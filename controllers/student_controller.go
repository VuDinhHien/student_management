package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"student_management/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	students, _ := models.GetAllStudents()
	tmpl.ExecuteTemplate(w, "index.html", students)
}

func ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "create.html", nil)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		email := r.FormValue("email")
		student := models.Student{Name: name, Age: age, Email: email}
		models.CreateStudent(student)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ShowEditForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	student, _ := models.GetStudentByID(id)
	tmpl.ExecuteTemplate(w, "edit.html", student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		email := r.FormValue("email")
		student := models.Student{ID: id, Name: name, Age: age, Email: email}
		models.UpdateStudent(student)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.DeleteStudent(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
