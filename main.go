package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/desertbit/fillpdf"
)

func main() {
	file, _ := os.Open("data.json")
	decoder := json.NewDecoder(file)
	data := new(DataToFill)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
		return
	}

	form := make(map[string]interface{})
	form, err = data.CreateForm(form)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = fillpdf.Fill(form, "f8949.pdf", "f8949_filled.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(form)
}

type DataToFill struct {
	Page1 Page
	Page2 Page
}

type Page struct {
	PageNumber int
	Name       string
	ID         string
	Box        int
	Rows       []Row
}

type Row struct {
	Columns []string
}

func (d DataToFill) CreateForm(form map[string]interface{}) (map[string]interface{}, error) {
	form, err := d.Page1.CreatePageForm(form)
	if err != nil {
		return nil, err
	}
	form, err = d.Page2.CreatePageForm(form)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func (p Page) CreatePageForm(form map[string]interface{}) (map[string]interface{}, error) {
	if p.Box < 1 || p.Box > 3 {
		return nil, fmt.Errorf("the box parameter must take values from 1 to 3")
	}
	if len(p.Rows) > 14 {
		return nil, fmt.Errorf("rows cannot be more than 14")
	}
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_1[0]", p.PageNumber, p.PageNumber)] = p.Name
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_2[0]", p.PageNumber, p.PageNumber)] = p.ID
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].c%d_1[%d]", p.PageNumber, p.PageNumber, p.Box-1)] = p.Box

	id := 3
	for i := range p.Rows {
		if len(p.Rows[i].Columns) > 8 {
			return nil, fmt.Errorf("columns cannot be more than 8")
		}
		for j := 0; j < 8; j++ {
			if len(p.Rows[i].Columns) > j {
				form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", p.PageNumber, i+1, p.PageNumber, id)] = p.Rows[i].Columns[j]
			}
			id++
		}
	}
	return form, nil
}
