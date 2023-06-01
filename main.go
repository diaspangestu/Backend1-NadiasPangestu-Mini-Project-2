package main

import (
	"fmt"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/db"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/customer"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// open connection db
	db := db.InitDB()

	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
