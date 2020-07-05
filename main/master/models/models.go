package models

type Product struct {
	Id          string `json:"Id"`
	ProductCode string `json:"ProductCode"`
	ProductName string `json:"ProductName"`
	CategoryID  string `json:"CategoryID"`
	Price       string `json:"Price"`
	CreatedAt   string `json:"Created"`
	UpdatedAt   string `json:"Updated"`
}

type Category struct {
	Id           string `json:"Id"`
	CategoryName string `json:"CategoryName"`
	CreatedAt    string `json:"Created"`
	UpdatedAt    string `json:"Updated"`
}

type Order struct {
	IdProduct string `json:"IdProduct"`
	OrderDate string `json:"OrderDate"`
	CreatedAt string `json:"Created"`
	UpdatedAt string `json:"Updated"`
}

type OrderItem struct {
	OrderId   string `json:"OrderId"`
	Qty       string `json:"Quantity"`
	ProductId string `json:"CategoryID"`
	CreatedAt string `json:"Created"`
	UpdatedAt string `json:"Updated"`
}

type DailyReport struct {
	Day         string `json:"Day"`
	Month       string `json:"Month"`
	Year        string `json:"Year"`
	ProductName string `json:"Productname"`
	Price       string `json:"Price"`
	Qty         string `json:"Quantity"`
	Total       string `json:"Total"`
}

type MonthlyReport struct {
	Month string `json:"Month"`
	Year  string `json:"Year"`
	Qty   string `json:"ItemSold"`
	Total string `json:"TotalProfit"`
}
