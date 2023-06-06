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

func (rh RequestHandlerAdmin) LoginAdmin(c *gin.Context) {
	request := AdminParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.LoginAdmin(request.ID, request.Username, request.Password)
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

func (rh RequestHandlerAdmin) GetAllCustomers(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	email := c.Query("email")

	// Parse pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	res, err := rh.ctrl.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}
