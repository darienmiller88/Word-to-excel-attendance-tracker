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

	doc, err := document.Open("./docs/Pelham/October/10_13_25 Pelham.docx")

	if err != nil {
		panic(err)
	}

	defer doc.Close()

	attendance := Attendance{}
	// attendances := []Attendance{}
	addedDate, addedTime, addedLocation, addedCourse := false, false, false, false

	for _, paragraph := range doc.Paragraphs() {
		for _, r := range paragraph.Runs() {

			if strings.HasPrefix(r.Text(), "Date"){
				attendance.date = strings.Split(r.Text(), " ")[1]
				addedDate = true
			}

			if strings.HasPrefix(r.Text(), "Location"){
				attendance.location = strings.Split(r.Text(), " ")[1]
				addedLocation = true
			}

			if strings.HasPrefix(r.Text(), "Course") {
				course, _ := strings.CutPrefix(r.Text(), "Course: ")
				attendance.course = course
				addedCourse = true
			}

			if strings.HasPrefix(r.Text(), "Time") {
				time, _ := strings.CutPrefix(r.Text(), "Time: ")
				attendance.time = time
				addedTime = true
			}

			if addedCourse && addedDate && addedLocation && addedTime {
				fmt.Println(attendance)				
				fmt.Println()
				addedCourse, addedDate, addedTime, addedLocation = false, false, false, false
			}
		}
	}
}
