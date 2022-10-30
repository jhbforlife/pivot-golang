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

	initProducts()

	r := mux.NewRouter()
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", AddProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", GetProductIDHandler).Methods("GET")
	r.HandleFunc("/products/{id}", UpdateProductIDHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProductIDHandler).Methods("DELETE")
	r.PathPrefix("/").HandlerFunc(CatchAllHandler)

	log.Println("Starting server. Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func initProducts() {
	log.Println("Initializing products. Reading JSON file...")
	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read successful. Unmarshaling JSON...")
	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
	log.Println("Unmarshal successful. Products initialized.")
}

func checkInternalError(err error, w http.ResponseWriter) (wasError bool) {
	if err != nil {
		log.Println(err)
		writeHTTPStatus(http.StatusInternalServerError, w)
		return true
	} else {
		return false
	}
}

func writeToFile(bs []byte, w http.ResponseWriter) (didWrite bool) {
	err := os.WriteFile("products.json", bs, 0666)
	return !checkInternalError(err, w)
}

func writeHTTPStatus(status int, w http.ResponseWriter) {
	statusWithText := strconv.Itoa(status) + " " + http.StatusText(status) + "\n"
	w.WriteHeader(status)
	w.Write([]byte(statusWithText))
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received product GET request. Marshaling all product data...")
	bs, err := json.Marshal(products)
	if checkInternalError(err, w) {
		return
	}
	log.Println("Marshal successful. Writing JSON...")
	w.Header().Add("content-type", "application/json")
	if _, err := w.Write(bs); checkInternalError(err, w) {
		return
	}
	log.Println("Write successful.")
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct product
	var newProducts []product
	log.Println("Received product POST request. Reading incoming JSON...")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if checkInternalError(err, w) {
		return
	}
	log.Println("Read successful. Unmarshaling product...")
	if err := json.Unmarshal(body, &newProduct); checkInternalError(err, w) {
		return
	}
	log.Println("Unmarshal successful. Adding to existing products...")
	newProducts = append(products, newProduct)
	bs, err := json.Marshal(newProducts)
	if checkInternalError(err, w) {
		return
	}
	if !writeToFile(bs, w) {
		return
	}
	products = newProducts
	log.Println(products)
	writeHTTPStatus(http.StatusCreated, w)
	log.Println("Successfully added product to list.")
}

func GetProductIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Received product GET request for ID", vars["id"], "- checking for product...")
	for _, v := range products {
		if strconv.Itoa(v.ID) == vars["id"] {
			log.Println("Product ID found. Marshaling product data...")
			bs, err := json.Marshal(v)
			if checkInternalError(err, w) {
				return
			}
			w.Header().Add("content-type", "application/json")
			log.Println("Marshal successful. Writing JSON...")
			if _, err := w.Write(bs); checkInternalError(err, w) {
				return
			}
			log.Println("Write successful.")
			return
		}
	}
	writeHTTPStatus(http.StatusNotFound, w)
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
			if checkInternalError(err, w) {
				return
			}
			defer r.Body.Close()
			log.Println("Read successful. Unmarshaling data...")
			if err := json.Unmarshal(bs, &newProductInfo); checkInternalError(err, w) {
				return
			}
			log.Println("Unmarshal successful. Marshaling new product data...")
			if name, ok := newProductInfo["name"]; ok {
				v.Name = name.(string)
			} else {
				log.Println("New name not found...")
			}
			if price, ok := newProductInfo["price"]; ok {
				v.Price = int(price.(float64))
			} else {
				log.Println("New id not found...")
			}
			updatedProduct = true
		}
		newProducts = append(newProducts, v)
	}
	if updatedProduct {
		ns, err := json.Marshal(newProducts)
		if checkInternalError(err, w) {
			return
		}
		log.Println("Marshal successful. Writing new product data...")
		if !writeToFile(ns, w) {
			return
		}
		products = newProducts
		writeHTTPStatus(http.StatusAccepted, w)
		log.Println("Write successful. Product info updated.")
	} else {
		writeHTTPStatus(http.StatusNotFound, w)
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
		bs, err := json.Marshal(newProducts)
		if checkInternalError(err, w) {
			return
		}
		log.Println("Marshal successful. Writing new product data...")
		if !writeToFile(bs, w) {
			return
		}
		products = newProducts
		writeHTTPStatus(http.StatusAccepted, w)
		log.Println("Write successful. Product info updated.")
	} else {
		writeHTTPStatus(http.StatusNotFound, w)
		log.Println("Product ID not found...")
	}
}

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("URL", r.URL, "requested. Not a valid URL...")
	writeHTTPStatus(http.StatusNotFound, w)
}
