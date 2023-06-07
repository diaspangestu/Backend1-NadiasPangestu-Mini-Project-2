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

func (rh RequestHandlerSuperadmin) CreateSuperadmin(c *gin.Context) {
	request := SuperAdminParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.CreateSuperadmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandlerSuperadmin) LoginSuperadmin(c *gin.Context) {
	request := SuperAdminParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.LoginSuperadmin(request.ID, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
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

// DeleteCustomerById Superadmin
func (rh RequestHandlerSuperadmin) DeleteCustomerById(c *gin.Context) {
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

func (rh RequestHandlerSuperadmin) GetAllCustomers(c *gin.Context) {
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

func (rh RequestHandlerSuperadmin) UpdateActivedAdmin(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	activeID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.UpdateActivedAdmin(uint(activeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerSuperadmin) UpdateDeadactivedAdmin(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	activeID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.UpdateDeadactivedAdmin(uint(activeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerSuperadmin) GetApprovalRequest(c *gin.Context) {
	response, err := rh.ctrl.GetApprovalRequest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerSuperadmin) GetAllAdmins(c *gin.Context) {
	username := c.Query("username")

	// Parse pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	res, err := rh.ctrl.GetAllAdmins(username, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}
