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

type OrderItemHandler struct {
	OrderItemUseCase usecases.OrderItemUseCase
}

func OrderItemController(r *mux.Router, service usecases.OrderItemUseCase) {
	OrderItemHandler := OrderItemHandler{service}
	r.HandleFunc("/orderitems", OrderItemHandler.ListOrderItems).Methods(http.MethodGet)
	r.HandleFunc("/orderitem/{OrderId}", OrderItemHandler.GetOrderItemID).Methods(http.MethodGet)
	r.HandleFunc("/orderitem", OrderItemHandler.CreateDataOrderItems).Methods(http.MethodPost)
	r.HandleFunc("/orderitem", OrderItemHandler.UpdateDataOrderItems).Methods(http.MethodPut)
	r.HandleFunc("/orderitem/{OrderId}", OrderItemHandler.DeleteDataOrderItems).Methods(http.MethodDelete)
}

func (s OrderItemHandler) ListOrderItems(w http.ResponseWriter, r *http.Request) {
	orderitem, err := s.OrderItemUseCase.GetAllOrderItems()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database List Order Items"
	pesan.Data = orderitem

	byteOfOrderItems, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrderItems)
	fmt.Println("EndPoint Hit : GetOrderItemsPage")
}

func (s OrderItemHandler) GetOrderItemID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderitem, err := s.OrderItemUseCase.GetOrderItemByID(vars["OrderId"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database List Order Items by ID"
	pesan.Data = orderitem

	byteOfOrderItems, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrderItems)
	fmt.Println("EndPoint Hit : GetOrderItemsByIDPage")

}

func (s OrderItemHandler) CreateDataOrderItems(w http.ResponseWriter, r *http.Request) {
	var orderitem models.OrderItem
	_ = json.NewDecoder(r.Body).Decode(&orderitem) // json ke struct
	s.OrderItemUseCase.CreateOrderItems(orderitem)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Data for Order Item Database"
	pesan.Data = "Success"

	byteOfOrderItems, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrderItems)
	fmt.Println("EndPoint Hit : CreateOrderItemitemPage")
}

func (s OrderItemHandler) UpdateDataOrderItems(w http.ResponseWriter, r *http.Request) {
	var orderitem models.OrderItem
	_ = json.NewDecoder(r.Body).Decode(&orderitem) // json ke struct
	s.OrderItemUseCase.UpdateOrderItems(orderitem)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Data for Order Item Database"
	pesan.Data = "Success"

	byteOfOrderItems, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrderItems)
	fmt.Println("EndPoint Hit : UpdateOrderItemsPage")

}

func (s OrderItemHandler) DeleteDataOrderItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.OrderItemUseCase.DeleteOrderItems(vars["OrderId"])

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Data for Order Item Database"
	pesan.Data = "Success"

	byteOfOrderItems, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfOrderItems)
	fmt.Println("EndPoint Hit : DeleteOrderItemsPage")
}
