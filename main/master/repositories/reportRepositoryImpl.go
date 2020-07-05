package repositories

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/models"
)

type ReportRepoImpl struct {
	db *sql.DB
}

func (s ReportRepoImpl) GetAllDailies() ([]*models.DailyReport, error) {
	dataDaily := []*models.DailyReport{}
	query := "select day(o.order_date), month(o.order_date), year(o.order_date), p.product_name, p.price, (poi.qty), (p.price*poi.qty) from purchase_order o join purchase_order_item poi on poi.product_id=o.id_product join product p on poi.product_id=p.id group by 4, 3, 2, 1, 5 order by 2"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Daily := models.DailyReport{}
		var err = data.Scan(&Daily.Day, &Daily.Month, &Daily.Year, &Daily.ProductName, &Daily.Price, &Daily.Qty, &Daily.Total)
		if err != nil {
			return nil, err
		}
		dataDaily = append(dataDaily, &Daily)
	}
	return dataDaily, nil
}
func (s ReportRepoImpl) GetAllMonthlies() ([]*models.MonthlyReport, error) {
	dataMonthly := []*models.MonthlyReport{}
	query := "select month(o.order_date), year(o.order_date), count(poi.qty), sum(p.price*poi.qty) from purchase_order o join purchase_order_item poi on poi.product_id=o.id_product join product p on o.id_product=p.id group by 1"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Monthly := models.MonthlyReport{}
		var err = data.Scan(&Monthly.Month, &Monthly.Year, &Monthly.Qty, &Monthly.Total)
		if err != nil {
			return nil, err
		}
		dataMonthly = append(dataMonthly, &Monthly)
	}
	return dataMonthly, nil
}

func InitReportRepoImpl(db *sql.DB) ReportRepository {
	return &ReportRepoImpl{db}
}
