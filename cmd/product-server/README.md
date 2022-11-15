# Assignment: Build a Product Server

## Server Instructions and Examples of Implementation
1. `GET /products` returns all products  

	input:  
	`curl localhost:8080/products`  
	
	output:  
	`[{"id":1,"name":"Water - San Pellegrino","description":"curae nulla dapibus dolor vel est donec odio justo sollicitudin ut suscipit a feugiat et eros vestibulum ac est lacinia","price":80},...{"id":100,"name":"Shrimp - Black Tiger 16/20","description":"ultrices enim lorem ipsum dolor sit amet consectetuer adipiscing elit proin","price":63}]`
  ##
2. `GET /products/{id}` returns the specified product or a 404 status code if not found

	input:  
	`curl localhost:8080/products/60`
	
	output:  
	`{"id":60,"name":"Cookies - Englishbay Wht","description":"rutrum nulla nunc purus phasellus in felis donec semper sapien a libero nam dui proin leo odio porttitor id","price":46}`

	input:  
	`curl localhost:8080/products/200`
	
	output:  
	`404 Not Found`
##
3. `POST /products` adds a new product to the server which should be listed on subsequent calls to retrieve all products

	input:  
	`curl -X POST localhost:8080/products -H "content-type: application/json" -d '{"id": 101, "name": "AirPods Pro (2nd gen)", "description":"Newest earbuds from Apple", "price": 249}'`
	
	output:  
	`201 Created`
##	
4. `PUT /products/{id}` updates a product’s name and price

	input:  
	`curl -X PUT localhost:8080/products/101 -H "content-type: application/json" -d '{"name": "AirPods Pro", "price": 199}'`
	
	output:  
	`202 Accepted`
##	
5. `DELETE /products/{id}` removes a product from the server

	input:  
	`curl -X DELETE localhost:8080/products/101`
	
	output:  
	`202 Accepted`

## Assignment Instructions  
### Objective

Build an HTTP server that supports returning product information such as name and price. The server should also support the ability to retrieve the complete list of products or a single product based on its ID. Adding and updating new products are also expected.

### Learning Goals

1. Understand the use of the [net/http](https://pkg.go.dev/net/http) and [encoding/json](https://pkg.go.dev/encoding/json) standard library packages
2. Leverage third-party libraries like [gorilla/mux](https://github.com/gorilla/mux) 
3. Work with JSON data (read from file, unmarshal into a Go struct)
4. Practice using GitHub to house your projects

### Acceptance Criteria

1. Product data is externalized in a JSON file (products.json)
2. When server starts, it reads and unmarshals the data into data structure suitable for lookups
    1. Use [JSON-to-Go](https://mholt.github.io/json-to-go/) to convert a product JSON sample into a Go struct
3. Server endpoints:
    1. `GET /products` returns all products
    2. `GET /products/{id}` returns the specified product or a 404 status code if not found
    3. `POST /products` adds a new product to the server which should be listed on subsequent calls to retrieve all products
    4. `PUT /products/{id}` updates a product’s name and price
    5. `DELETE /products/{id}` removes a product from the server
4. Provide a link to a GitHub pull request (PR) against your repository’s main branch with your solution for evaluation.

### Resources

- [Instructor Video - Building an HTTP Server](https://drive.google.com/file/d/1cF6MNqliUzYUvqbliz7j1QRx3y4wx749/view?usp=sharing)
- [Sample products dataset](https://gist.githubusercontent.com/jboursiquot/259b83a2d9aa6d8f16eb8f18c67f5581/raw/9b28998704fb06f127f13540a4f6e3812f50774b/products.json)

### Notes

Product additions, updates, and deletions are not expected to be reflected in the products.json file you initialize the server with.
(I did not see this, thus, my server reflects changes in my products.json file.)
