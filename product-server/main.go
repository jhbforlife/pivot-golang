package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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

	printInstructions()

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

func printInstructions() {
	log.Print(`
Instructions for product-server:
1 - GET /products returns all products:

curl localhost:8080/products

2 - GET /products/{id} returns the specified product or a 404 status code if not found:

curl localhost:8080/products/60
curl localhost:8080/products/200

3 - POST /products adds a new product to the server which should be listed on subsequent calls to retrieve all products:

curl -X POST localhost:8080/products -H "content-type: application/json" -d {"id": 101, "name": "AirPods Pro (2nd gen)", "description":"Newest earbuds from Apple", "price": 249}
	
4 - PUT /products/{id} updates a product's name and price:

curl -X PUT localhost:8080/products/101 -H "content-type: application/json" -d {"name": "AirPods Pro", "price": 199}
	
5 - DELETE /products/{id} removes a product from the server:

curl -X DELETE localhost:8080/products/101

`)
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
	bs, err := json.MarshalIndent(products, "", "  ")
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
	log.Println("Write successful.")
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct product
	log.Println("Received product POST request. Reading incoming JSON...")
	body, err := io.ReadAll(r.Body)
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
	bs, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := os.WriteFile("products.json", bs, 0666); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("Successfully added product to list.")
}

func GetProductIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Received product GET request for ID", vars["id"], "- checking for product...")
	for _, v := range products {
		if strconv.Itoa(v.ID) == vars["id"] {
			log.Println("Product ID found. Marshaling product data...")
			bs, err := json.MarshalIndent(v, "", "")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			w.Header().Add("content-type", "application/json")
			log.Println("Marshal successful. Writing JSON...")
			if _, err := w.Write(bs); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			log.Println("Write successful.")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	log.Println("Product ID not found...")
}

func UpdateProductIDHandler(w http.ResponseWriter, r *http.Request) {
	var newProductInfo map[string]interface{}
	var updatedProduct bool
	newProducts := []product{}
	vars := mux.Vars(r)
	log.Println("Received product PUT request for ID", vars["id"], "- checking for product...")
	for _, v := range products {
		if strconv.Itoa(v.ID) == vars["id"] {
			log.Println("Product ID found. Reading new product data...")
			bs, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			defer r.Body.Close()
			log.Println("Read successful. Unmarshaling data...")
			if err := json.Unmarshal(bs, &newProductInfo); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			log.Println("Unmarshal successful. Marshaling new product data...")
			if name, ok := newProductInfo["name"]; ok {
				v.Name = name.(string)
			} else {
				log.Println("New name not found...")
			}
			if price, ok := newProductInfo["price"]; ok {
				v.ID = int(price.(float64))
			} else {
				log.Println("New id not found...")
			}
			updatedProduct = true
		}
		newProducts = append(newProducts, v)
	}
	if updatedProduct {
		products = newProducts
		ns, err := json.MarshalIndent(products, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		log.Println("Marshal successful. Writing new product data...")
		if err := os.WriteFile("products.json", ns, 0666); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		log.Println("Write successful. Product info updated.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Product ID not found...")
	}
}

func DeleteProductIDHandler(w http.ResponseWriter, r *http.Request) {
	var deletedProduct bool
	newProducts := []product{}
	vars := mux.Vars(r)
	log.Println("Received product DELETE request for ID", vars["id"], "- checking for product...")
	for _, v := range products {
		if strconv.Itoa(v.ID) != vars["id"] {
			newProducts = append(newProducts, v)
		} else {
			deletedProduct = true
		}
	}
	if deletedProduct {
		log.Println("Product ID found. Marshaling new product data...")
		products = newProducts
		bs, err := json.MarshalIndent(products, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		log.Println("Marshal successful. Writing new product data...")
		if err := os.WriteFile("products.json", bs, 0666); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		log.Println("Write successful. Product info updated.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Product ID not found...")
	}
}
