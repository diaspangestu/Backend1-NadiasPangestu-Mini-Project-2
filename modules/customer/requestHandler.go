package customer

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	ctrl ControllerCustomer
}

func NewCustomerRequestHandler(
	db *gorm.DB,
) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctrl: ControllerCustomer{
			uc: UsecaseCustomer{
				customerRepo: repositories.NewCustomer(db),
			},
		},
	}
}

func (rh RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	req := CustomerParam{}

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	res, err := rh.ctrl.CreateCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.GetCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}
