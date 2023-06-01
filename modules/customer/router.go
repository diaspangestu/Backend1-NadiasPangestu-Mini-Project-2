package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewRouter(
	db *gorm.DB,
) RouterCustomer {
	return RouterCustomer{CustomerRequestHandler: NewCustomerRequestHandler(db)}
}

func (r RouterCustomer) Handle(router *gin.Engine) {
	basePath := "/"

	customer := router.Group(basePath)
	customer.POST("customer", r.CustomerRequestHandler.CreateCustomer)
	customer.GET("customer/:id", r.CustomerRequestHandler.GetCustomerById)
	customer.PUT("customer/:id", r.CustomerRequestHandler.UpdateCustomerById)
	customer.DELETE("customer/:id", r.CustomerRequestHandler.DeleteCustomerById)
}
