package main

import (
	"ecommerce-crud-middleware/config"
	"ecommerce-crud-middleware/main/master"
)

func main() {
	db, _ := config.Connection()    //ok
	router := config.CreateRouter() //ok
	master.Init(router, db)
	config.RunServer(router)
}
