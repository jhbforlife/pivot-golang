package main

import (
	"encoding/json"
	"io/ioutil"
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
	log.Println("Received product GET request. Marshaling all product data...")
	bs, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("Marshal successful. Writing JSON...")
	w.Header().Add("content-type", "application/json")
	if _, err := w.Write(bs); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Write successful.")
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct product
	log.Println("Received product POST request. Reading incoming JSON...")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("Read successful. Unmarshaling product...")
	if err := json.Unmarshal(body, &newProduct); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("Unmarshal successful. Adding to existing products...")
	products = append(products, newProduct)
	bs, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := os.WriteFile("products.json", bs, 0666); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Successfully added product to list.")
}

func GetProductIDHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateProductIDHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteProductIDHandler(w http.ResponseWriter, r *http.Request) {

}
