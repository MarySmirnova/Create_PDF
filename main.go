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
	r.HandleFunc("/fillpdf", completeDocument)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func completeDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "wrong method", 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data := new(form.DataToFill)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	form, err := data.CreateForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	err = fillpdf.Fill(form, "f8949.pdf", fmt.Sprintf("data/%s.pdf", data.Pages[0].Name))
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("create file %s.pdf", data.Pages[0].Name)
	w.Write([]byte("filling pdf file completed"))
}
