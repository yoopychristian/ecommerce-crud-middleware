package usecases

import "ecommerce-crud-middleware/main/master/models"

type ReportUseCase interface {
	GetAllDailies() ([]*models.DailyReport, error)
	GetAllMonthlies() ([]*models.MonthlyReport, error)
}
