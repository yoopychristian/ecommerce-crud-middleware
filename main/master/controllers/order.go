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

type OrderHandler struct {
	OrderUseCase usecases.OrderUseCase
}

func OrderController(r *mux.Router, service usecases.OrderUseCase) {
	OrderHandler := OrderHandler{service}
	r.HandleFunc("/orders", OrderHandler.ListOrders).Methods(http.MethodGet)
	r.HandleFunc("/order/{IdProduct}", OrderHandler.GetOrderID).Methods(http.MethodGet)
	r.HandleFunc("/order", OrderHandler.CreateDataOrders).Methods(http.MethodPost)
	r.HandleFunc("/order", OrderHandler.UpdateDataOrders).Methods(http.MethodPut)
	r.HandleFunc("/order/{IdProduct}", OrderHandler.DeleteDataOrders).Methods(http.MethodDelete)
}

func (s OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	order, err := s.OrderUseCase.GetAllOrders()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database List Orders"
	pesan.Data = order

	byteOfOrders, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrders)
	fmt.Println("EndPoint Hit : GetOrderPage")
}

func (s OrderHandler) GetOrderID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order, err := s.OrderUseCase.GetOrderByID(vars["IdProduct"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database List Orders By ID"
	pesan.Data = order

	byteOfOrders, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrders)
	fmt.Println("EndPoint Hit : GetOrderPage By ID")
}

func (s OrderHandler) CreateDataOrders(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	_ = json.NewDecoder(r.Body).Decode(&order) // json ke struct
	s.OrderUseCase.CreateOrders(order)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Data for Orders Database"
	pesan.Data = "Success"

	byteOfOrders, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfOrders)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateOrderPage")
}

func (s OrderHandler) UpdateDataOrders(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	_ = json.NewDecoder(r.Body).Decode(&order) // json ke struct
	s.OrderUseCase.UpdateOrders(order)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Data for Order Database"
	pesan.Data = "Success"

	byteOfOrders, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfOrders)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateOrderPage")
}

func (s OrderHandler) DeleteDataOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.OrderUseCase.DeleteOrders(vars["IdProduct"])

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Data for Order Database"
	pesan.Data = "Success"

	byteOfOrders, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfOrders)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteOrderPage")

}
