package master

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/controllers"
	"ecommerce-crud-middleware/main/master/repositories"
	"ecommerce-crud-middleware/main/master/usecases"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	productRepo := repositories.InitProductRepoImpl(db)
	productUseCase := usecases.InitProductUseCase(productRepo)
	controllers.ProductController(r, productUseCase)
	categoryRepo := repositories.InitCategoryRepoImpl(db)
	categoryUseCase := usecases.InitCategoryUseCase(categoryRepo)
	controllers.CategoryController(r, categoryUseCase)
	orderRepo := repositories.InitOrderRepoImpl(db)
	orderUseCase := usecases.InitOrderUseCase(orderRepo)
	controllers.OrderController(r, orderUseCase)
	orderItemRepo := repositories.InitOrderItemRepoImpl(db)
	orderItemUseCase := usecases.InitOrderItemUseCase(orderItemRepo)
	controllers.OrderItemController(r, orderItemUseCase)
	reportRepo := repositories.InitReportRepoImpl(db)
	reportUseCase := usecases.InitReportUseCase(reportRepo)
	controllers.ReportController(r, reportUseCase)
}
