package middleware

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
