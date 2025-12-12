package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
)

func main(){
	godotenv.Load()

	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	
	if err != nil {
		panic(err)
	}

	doc, err := document.Open("./docs/Pelham/October/10_13_25 Pelham.docx")

	if err != nil {
		panic(err)
	}

	fmt.Println("notes:", doc.ExtractText().Text())
}