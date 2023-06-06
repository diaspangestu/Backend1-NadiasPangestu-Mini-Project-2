package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterAdmin struct {
	AdminRequestHandler RequestHandlerAdmin
}

func NewRouter(db *gorm.DB) RouterAdmin {
	return RouterAdmin{AdminRequestHandler: NewAdminRequestHandler(db)}
}

func (r RouterAdmin) Handle(router *gin.Engine) {
	basePath := "/admin"

	admin := router.Group(basePath)
	admin.POST("/login", r.AdminRequestHandler.LoginAdmin)
	admin.POST("/register-admin", r.AdminRequestHandler.RegisterAdmin)

	// About Customer
	admin.Use(middleware.Authentication())
	admin.POST("/create-customer", r.AdminRequestHandler.CreateCustomer)
	admin.DELETE("/delete-customer/:id", r.AdminRequestHandler.DeleteCustomerById)
	admin.GET("/customers", r.AdminRequestHandler.GetAllCustomers)
}
