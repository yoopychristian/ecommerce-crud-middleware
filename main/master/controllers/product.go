package controllers

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/response"
	"ecommerce-crud-middleware/main/master/usecases"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductUseCase usecases.ProductUseCase
}

func ProductController(r *mux.Router, service usecases.ProductUseCase) {
	ProductHandler := ProductHandler{service}
	r.HandleFunc("/products", ProductHandler.ListProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{Id}", ProductHandler.GetProductID).Methods(http.MethodGet)
	r.HandleFunc("/product", ProductHandler.CreateDataProducts).Methods(http.MethodPost)
	r.HandleFunc("/product", ProductHandler.UpdateDataProducts).Methods(http.MethodPut)
	r.HandleFunc("/product/{Id}", ProductHandler.DeleteDataProducts).Methods(http.MethodDelete)
}

func (s ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	product, err := s.ProductUseCase.GetProducts()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database All Products"
	pesan.Data = product

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfProducts)
	fmt.Println("Endpoint Hit: GetProductPage")
}

func (s ProductHandler) GetProductID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product, err := s.ProductUseCase.GetProductByID(vars["Id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Product by ID"
	pesan.Data = product

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfProducts)
	fmt.Println("Endpoint Hit: GetProductByIDPage")
}

func (s ProductHandler) CreateDataProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product) // json ke struct
	s.ProductUseCase.CreateProducts(product)
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Database Product"
	pesan.Data = "Success"

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfProducts)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateProductPage")
}

func (s ProductHandler) UpdateDataProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product) // json ke struct
	s.ProductUseCase.UpdateProducts(product)
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Database Product"
	pesan.Data = "Success"

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfProducts)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateProductPage")
}

func (s ProductHandler) DeleteDataProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.ProductUseCase.DeleteProducts(vars["Id"])
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Database Product"
	pesan.Data = "Success"

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfProducts)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteProductPage")
}
