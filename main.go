package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
)

//Attendance object to store attedance data for each docx file.
type Attendance struct{
	date       string
	time       string
	course     string
	location   string
	students []string
}

// push from same laptop, fixed lol
func main() {
	godotenv.Load()

	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))

	if err != nil {
		panic(err)
	}

	doc, err := document.Open("./docs/Pelham/November/11_3_25 Pelham.docx")

	if err != nil {
		panic(err)
	}

	defer doc.Close()

	attendance := Attendance{}
	attendances := []Attendance{}
	addedDate, addedTime, addedLocation, addedCourse := false, false, false, false
	startAddingStudents := false

	for _, paragraph := range doc.Paragraphs() {
		for _, r := range paragraph.Runs() {

			if startAddingStudents && r.Text() != "" {
				attendance.students = append(attendance.students, r.Text())
			} else if startAddingStudents && len(r.Text()) == 0{
				startAddingStudents = false
			} else if strings.HasPrefix(r.Text(), "Date"){
				attendance.date = strings.Split(r.Text(), " ")[1]
				addedDate = true
			} else if strings.HasPrefix(r.Text(), "Location"){
				attendance.location = strings.Split(r.Text(), " ")[1]
				addedLocation = true
			} else if strings.HasPrefix(r.Text(), "Course") {
				course, _ := strings.CutPrefix(r.Text(), "Course: ")
				attendance.course = course
				addedCourse = true
			} else if strings.HasPrefix(r.Text(), "Time") {
				time, _ := strings.CutPrefix(r.Text(), "Time: ")
				attendance.time = time
				addedTime = true
			} else if strings.ToLower(r.Text()) == "participants:" {
				startAddingStudents = true
			}

			if addedCourse && addedDate && addedLocation && addedTime {
				attendances = append(attendances, attendance)
				addedCourse, addedDate, addedTime, addedLocation = false, false, false, false
			}
		}
	}

	fmt.Println(attendances)
}
