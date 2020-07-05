package controllers

import (
	"ecommerce-crud-middleware/main/master/middleware"
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/response"
	"ecommerce-crud-middleware/main/master/usecases"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	CategoryUseCase usecases.CategoryUseCase
}

func CategoryController(r *mux.Router, service usecases.CategoryUseCase) {
	r.Use(middleware.ActivityLogMiddleware)
	CategoryHandler := CategoryHandler{service}
	r.HandleFunc("/categories", CategoryHandler.ListCategories).Methods(http.MethodGet)
	r.HandleFunc("/category/{Id}", CategoryHandler.GetCategoryID).Methods(http.MethodGet)
	r.HandleFunc("/category", CategoryHandler.CreateDataCategories).Methods(http.MethodPost)
	r.HandleFunc("/category", CategoryHandler.UpdateDataCategories).Methods(http.MethodPut)
	r.HandleFunc("/category/{Id}", CategoryHandler.DeleteDataCategories).Methods(http.MethodDelete)
}

func (s CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	category, err := s.CategoryUseCase.GetAllCategories()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Categories"
	pesan.Data = category

	byteOfCategories, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfCategories)
	fmt.Println("EndPoint Hit : GetCategoryPage")
}

func (s CategoryHandler) GetCategoryID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category, err := s.CategoryUseCase.GetCategoryByID(vars["Id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Category by ID"
	pesan.Data = category

	byteOfCategories, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategories)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: GetCategoryHByIDPage")
}

func (s CategoryHandler) CreateDataCategories(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category) // json ke struct
	s.CategoryUseCase.CreateCategories(category)
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Data for Category Database"
	pesan.Data = "Success"

	byteOfCategories, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategories)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateCategoryPage")
}

func (s CategoryHandler) UpdateDataCategories(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category) // json ke struct
	s.CategoryUseCase.UpdateCategories(category)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Data for Category Database"
	pesan.Data = "Success"

	byteOfProducts, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfProducts)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateCategoryPage")
}

func (s CategoryHandler) DeleteDataCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.CategoryUseCase.DeleteCategories(vars["Id"])
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Database Category"
	pesan.Data = "Success"

	byteOfCategories, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategories)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteCategoryPage")
}
