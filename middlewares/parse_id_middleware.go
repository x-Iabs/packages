package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/x-Iabs/packages/types"
)

// ParseIDMiddlware is a middleware function that parses the "id" parameter from the request URL and sets it in the context.
// If the "id" parameter is not a valid integer, it returns a JSON response with a status code of 400 and an error message.
// The middleware function is designed to be used with the Gin framework.
func ParseIDMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response := types.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid ID",
				Errors:  []string{err.Error()},
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		c.Set("id", id)
		c.Next()
	}
}
