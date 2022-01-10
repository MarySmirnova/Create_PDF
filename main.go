package main

import (
	"fmt"

	"github.com/signintech/pdft"
)

func main() {
	/*
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Hello, world")
		err := pdf.OutputFileAndClose("hello.pdf")
		if err != nil {
			fmt.Println(err)
		}
	*/
	var pdf2 pdft.PDFt
	err := pdf2.Open("hello.pdf")
	if err != nil {
		fmt.Println("1", err)
	}
	fmt.Printf("%s", pdf2)

}
