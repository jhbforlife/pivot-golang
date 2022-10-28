package main

import "log"

func printInstructions() {
	log.Print(`
	Instructions for product-server:
	1 - GET /products returns all products
		curl localhost:8080/products
	2 - GET /products/{id} returns the specified product or a 404 status code if not found
		curl localhost:8080/products/60
		curl localhost:8080/products/200
	3 - POST /products adds a new product to the server which should be listed on subsequent calls to retrieve all products
		curl -X POST localhost:8080/products -H "content-type: application/json" -d {"id": 101, "name": "AirPods Pro (2nd gen)", "description":"Newest earbuds from Apple", "price": 249}
	4 - PUT /products/{id} updates a product\'s name and price
		curl -X PUT localhost:8080/products/101 -H "content-type: application/json" -d {"name": "AirPods Pro", "price": 199}
	5 - DELETE /products/{id} removes a product from the server
		curl -X DELETE localhost:8080/products/101
	`)
}
