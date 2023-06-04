package admin

import (
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
	admin.POST("/login", r.AdminRequestHandler.LoginSuperadmin)
	admin.POST("/create-customer", r.AdminRequestHandler.CreateCustomer)
	admin.POST("/register-admin", r.AdminRequestHandler.RegisterAdmin)
}
