package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	err = fillpdf.Fill(form, "f8949.pdf", "filled.pdf")
	if err != nil {
		log.Println(err)
		return
	}
	output, _ := os.ReadFile("filled.pdf")
	log.Println("create file")
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(201)
	w.Write(output)
	os.Remove("filled.pdf")
}
