package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerSuperadmin struct {
	ctrl ControllerSuperadmin
}

func NewSuperadminRequestHandler(db *gorm.DB) RequestHandlerSuperadmin {
	return RequestHandlerSuperadmin{
		ctrl: ControllerSuperadmin{
			uc: UsecaseSuperadmin{
				superadminRepo: repositories.NewSuperadmin(db),
			},
		},
	}
}

// CreateCustomer Superadmin
func (rh RequestHandlerSuperadmin) CreateCustomer(c *gin.Context) {
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

func (rh RequestHandlerSuperadmin) ApprovedAdminRegister(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	approvedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.ApprovedAdminRegister(uint(approvedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerSuperadmin) RejectedAdminRegister(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	rejectedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.RejectedAdminRegister(uint(rejectedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}
