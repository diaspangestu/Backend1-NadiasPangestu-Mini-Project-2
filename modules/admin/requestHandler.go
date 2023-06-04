package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (rh RequestHandlerAdmin) LoginSuperadmin(c *gin.Context) {
	request := AdminParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.LoginAdmin(request.Username, request.Password)
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

// DeleteCustomerById Admin
func (rh RequestHandlerAdmin) DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	err = rh.ctrl.DeleteCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete Customer Data Successfully",
	})
}
