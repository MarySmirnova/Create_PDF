package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MarySmirnova/create_pdf/form"

	"github.com/desertbit/fillpdf"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fillpdf", completeDocument).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func completeDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := new(form.DataToFill)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	form, err := data.CreateForm()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = fillpdf.Fill(form, "f8949.pdf", fmt.Sprintf("data/%s.pdf", data.Pages[0].Name))
	if err != nil {
		log.Fatal(err)
	}
}
