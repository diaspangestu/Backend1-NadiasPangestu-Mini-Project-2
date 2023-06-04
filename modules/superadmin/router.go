package superadmin

import (
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
	superadmin.POST("/create-customer", r.SuperadminRequestHandler.CreateCustomer)
	superadmin.POST("/:id/approved", r.SuperadminRequestHandler.ApprovedAdminRegister)
	superadmin.POST("/:id/rejected", r.SuperadminRequestHandler.RejectedAdminRegister)
	superadmin.GET("/approval-request", r.SuperadminRequestHandler.GetApprovalRequest)
}
