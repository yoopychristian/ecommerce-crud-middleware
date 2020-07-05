package controllers

import (
	"ecommerce-crud-middleware/main/master/response"
	"ecommerce-crud-middleware/main/master/usecases"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ReportHandler struct {
	ReportUseCase usecases.ReportUseCase
}

func ReportController(r *mux.Router, service usecases.ReportUseCase) {
	ReportHandler := ReportHandler{service}
	r.HandleFunc("/dailyreport", ReportHandler.ListDailies).Methods(http.MethodGet)
	r.HandleFunc("/monthlyreport", ReportHandler.ListMonthlies).Methods(http.MethodGet)
}

func (s ReportHandler) ListDailies(w http.ResponseWriter, r *http.Request) {
	report, err := s.ReportUseCase.GetAllDailies()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Daily Report"
	pesan.Data = report

	byteOfDailies, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfDailies)
	fmt.Println("EndPoint Hit : GetReportDailyPage")
}

func (s ReportHandler) ListMonthlies(w http.ResponseWriter, r *http.Request) {
	report, err := s.ReportUseCase.GetAllMonthlies()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Monthly Report"
	pesan.Data = report

	byteOfMonthlies, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfMonthlies)
	fmt.Println("EndPoint Hit : GetReportMonthlyPage")
}
