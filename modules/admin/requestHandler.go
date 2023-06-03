package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandlerAdmin struct {
	ctrl ControllerAdmin
}

func NewAdminRequestHandler(db *gorm.DB) RequestHandlerAdmin {
	return RequestHandlerAdmin{
		ctrl: ControllerAdmin{
			uc: UsecaseAdmin{
				adminRepo: repositories.NewAdmin(db),
			},
		},
	}
}

// CreateCustomer Admin
func (rh RequestHandlerAdmin) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandlerAdmin) RegisterAdmin(c *gin.Context) {
	request := AdminParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.RegisterAdmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}
