package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterSuperadmin struct {
	SuperadminRequestHandler RequestHandlerSuperadmin
}

func NewRouter(db *gorm.DB) RouterSuperadmin {
	return RouterSuperadmin{SuperadminRequestHandler: NewSuperadminRequestHandler(db)}
}

func (r RouterSuperadmin) Handle(router *gin.Engine) {
	basePath := "/superadmin"

	superadmin := router.Group(basePath)
	superadmin.POST("/register", r.SuperadminRequestHandler.CreateSuperadmin)
	superadmin.POST("/login", r.SuperadminRequestHandler.LoginSuperadmin)

	// About Data
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/customers", r.SuperadminRequestHandler.GetAllCustomers)
	superadmin.POST("/create-customer", r.SuperadminRequestHandler.CreateCustomer)
	superadmin.DELETE("/delete-customer/:id", r.SuperadminRequestHandler.DeleteCustomerById)

	// About Admin
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/admins", r.SuperadminRequestHandler.GetAllAdmins)
	superadmin.POST("/:id/approved", r.SuperadminRequestHandler.ApprovedAdminRegister)
	superadmin.POST("/:id/rejected", r.SuperadminRequestHandler.RejectedAdminRegister)
	superadmin.POST("/:id/actived", r.SuperadminRequestHandler.UpdateActivedAdmin)
	superadmin.POST("/:id/deadactived", r.SuperadminRequestHandler.UpdateDeadactivedAdmin)
	superadmin.GET("/approval-request", r.SuperadminRequestHandler.GetApprovalRequest)
}
