package repositories

import "ecommerce-crud-middleware/main/master/models"

type ReportRepository interface {
	GetAllDailies() ([]*models.DailyReport, error)
	GetAllMonthlies() ([]*models.MonthlyReport, error)
}
