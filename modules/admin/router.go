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
	basePath := "/"

	admin := router.Group(basePath)
	admin.POST("admin", r.AdminRequestHandler.RegisterAdmin)
}
