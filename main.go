package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
)

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

	for _, paragraph := range doc.Paragraphs() {
		for _, r := range paragraph.Runs() {
			if strings.HasPrefix(r.Text(), "Location") {
				fmt.Println(r.Text())
			}

			if strings.HasPrefix(r.Text(), "Course") {
				fmt.Println(r.Text())
			}

			if strings.HasPrefix(r.Text(), "Time") {
				fmt.Println(r.Text())
				fmt.Println()
			}
		}
	}
}
