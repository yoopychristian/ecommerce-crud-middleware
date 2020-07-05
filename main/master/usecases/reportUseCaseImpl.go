package usecases

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/repositories"
)

type ReportUsecaseImpl struct {
	reportRepo repositories.ReportRepository
}

func (s ReportUsecaseImpl) GetAllDailies() ([]*models.DailyReport, error) {
	daily, err := s.reportRepo.GetAllDailies()
	if err != nil {
		return nil, err
	}
	return daily, nil
}

func (s ReportUsecaseImpl) GetAllMonthlies() ([]*models.MonthlyReport, error) {
	monthly, err := s.reportRepo.GetAllMonthlies()
	if err != nil {
		return nil, err
	}
	return monthly, nil
}

func InitReportUseCase(reportRepo repositories.ReportRepository) ReportUseCase {
	return &ReportUsecaseImpl{reportRepo}
}
