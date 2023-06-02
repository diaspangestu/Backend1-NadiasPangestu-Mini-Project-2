package main

import (
	"fmt"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/db"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/admin"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/customer"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/superadmin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// open connection db
	db := db.InitDB()

	// Customer Handler
	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	// Admin Handler
	adminHandler := admin.NewRouter(db)
	adminHandler.Handle(router)

	// Superadmin Handler
	superadminHandler := superadmin.NewRouter(db)
	superadminHandler.Handle(router)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
