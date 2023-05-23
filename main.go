package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Key string `json:"key"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/post", handlePostRequest)

	fmt.Println("Server pokrenut na portu 8080...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Provera da li zahtev ima metodu POST
	if r.Method != "POST" {
		http.Error(w, "Dozvoljena je samo metoda POST.", http.StatusMethodNotAllowed)
		return
	}

	// Parsiranje JSON podataka iz tela zahteva
	var requestData RequestData

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Greška pri parsiranju JSON podataka.", http.StatusBadRequest)
		return
	}

	// Obrađivanje podataka
	response := ResponseData{
		Message: "Primljeni podaci: " + requestData.Key,
	}

	// Konverzija odgovora u JSON format
	jsonResponse, err := json.Marshal(response)

	fmt.Println(string(jsonResponse))

	if err != nil {
		http.Error(w, "Greška pri konverziji u JSON format.", http.StatusInternalServerError)
		return
	}

	// Postavljanje odgovarajućih HTTP zaglavlja
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Slanje odgovora nazad klijentu
	w.Write(jsonResponse)
}
