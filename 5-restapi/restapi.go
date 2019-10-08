package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

var products []Product

func productsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	unknownId := true
	for _, item := range products {
		if item.ID == params["id"] {
			unknownId = false
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	if unknownId {
		json.NewEncoder(w).Encode("error: ProductNotFound")
	} else {
		json.NewEncoder(w).Encode(&Product{})
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func putProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	unknownId := true
	for index, item := range products {
		if item.ID == params["id"] {
			unknownId = false
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	if unknownId {
		json.NewEncoder(w).Encode("error: ProductNotFound")
	} else {
		json.NewEncoder(w).Encode(products)
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	unknownId := true
	for index, item := range products {
		if item.ID == params["id"] {
			unknownId = false
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	if unknownId {
		json.NewEncoder(w).Encode("error: ProductNotFound")
	} else {
		json.NewEncoder(w).Encode(products)
	}
}

func main() {
	r := mux.NewRouter()
	products = append(products, Product{ID: "2121", Name: "Sony Playstation 4", Description: "Home video game console", Price: "500$"})
	products = append(products, Product{ID: "2122", Name: "Xbox One S", Description: "Home video game console", Price: "470$"})
	r.HandleFunc("/products_list", productsList).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", putProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
