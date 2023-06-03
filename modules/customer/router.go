package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewRouter(db *gorm.DB) RouterCustomer {
	return RouterCustomer{CustomerRequestHandler: NewCustomerRequestHandler(db)}
}

func (r RouterCustomer) Handle(router *gin.Engine) {
	basePath := "/customer"

	customer := router.Group(basePath)
	customer.POST("/create", r.CustomerRequestHandler.CreateCustomer)
	customer.GET("/:id", r.CustomerRequestHandler.GetCustomerById)
	customer.PUT("/:id", r.CustomerRequestHandler.UpdateCustomerById)
	customer.DELETE("/:id", r.CustomerRequestHandler.DeleteCustomerById)
}
