package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x-Iabs/packages/functions"
)

// ParsePaginationParams is a middleware function that parses the pagination parameters from the request.
// It extracts the offset and limit values from the request and sets them in the context for further processing.
// If there is an error while parsing the parameters, it handles the error and aborts the request with a bad request status.
func ParsePaginationParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		offset, limit, err := functions.ParsePaginationParams(c)
		if err != nil {
			functions.HandleErrorWithHttpStatus(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}

		c.Set("offset", offset)
		c.Set("limit", limit)
		c.Next()
	}
}
