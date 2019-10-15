package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func productsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product

	db := Database()
	db.Find(&products)

	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := params["id"]
	var product Product

	db := Database()
	db.First(&product, key)

	if product.ID == 0 {
		json.NewEncoder(w).Encode("error: ProductNotFound")
		return
	}
	fmt.Println("Endpoint Hit: Product #:",key)
	json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = rand.Intn(1000000)

	db := Database()
	db.Create(&product)

	fmt.Println("Endpoint Hit: Creating New Product, id:", product.ID)
	json.NewEncoder(w).Encode(product)
}

func putProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := params["id"]
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

    var dbProduct Product
	db := Database()
	db.First(&dbProduct, key)

	if dbProduct.ID == 0 {
		json.NewEncoder(w).Encode("error: ProductNotFound")
		return
	}

	db.Model(&dbProduct).Update("name", product.Name)
	db.Model(&dbProduct).Update("description", product.Description)
	db.Model(&dbProduct).Update("price", product.Price)
	fmt.Println("Endpoint Hit: Updating Product, id:", product.ID)
	json.NewEncoder(w).Encode(dbProduct)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)
	key := params["id"]

	db := Database()
	db.First(&product, key)

	if product.ID == 0 {
		json.NewEncoder(w).Encode("error: ProductNotFound")
		return
	}

	db.Delete(&product)
	fmt.Println("Endpoint Hit: Deleting Product, id:", product.ID)
	json.NewEncoder(w).Encode(product)
}

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "test:123@tcp(127.0.0.1:3306)/Shop?charset=utf8&parseTime=True")
	fmt.Println("Endpoint Hit: Connect db: ok")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	db := Database()
	db.AutoMigrate(&Product{})
	r := mux.NewRouter()
	r.HandleFunc("/products_list", productsList).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", putProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
