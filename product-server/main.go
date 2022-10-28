package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var products []product

func main() {

	initProducts()

	r := mux.NewRouter()
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", AddProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", GetProductIDHandler).Methods("GET")
	r.HandleFunc("/products/{id}", UpdateProductIDHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProductIDHandler).Methods("DELETE")

	log.Println("Starting server. Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func initProducts() {
	log.Println("Reading JSON file...")
	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read successful. Unmarshaling JSON...")
	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
	log.Println("Unmarshal successful.")
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {

}

func GetProductIDHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateProductIDHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteProductIDHandler(w http.ResponseWriter, r *http.Request) {

}
