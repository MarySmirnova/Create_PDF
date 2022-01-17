package main

import (
	"fmt"
	"os"

	"rsc.io/pdf"
)

func main() {

	file, err := pdf.Open("f8949.pdf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	page1 := file.Page(1)
	val := page1.V.TextFromUTF16()

	fmt.Println(val)
}
