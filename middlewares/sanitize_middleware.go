package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x-Iabs/packages/errors/messages"
	"github.com/x-Iabs/packages/sanitize"
	"github.com/x-Iabs/packages/types"
)

// SanitizeMiddleware is a middleware that sanitizes input to prevent XSS attacks
func SanitizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				response := types.ErrorResponse{
					Status:  http.StatusBadRequest,
					Message: messages.BadRequest,
					Errors:  []string{err.Error()},
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}

			// Sanitize the body
			sanitizedBody, err := sanitize.SanitizeJSON(body)
			if err != nil {
				response := types.ErrorResponse{
					Status:  http.StatusBadRequest,
					Message: messages.BadRequest,
					Errors:  []string{err.Error()},
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}

			// Replace the request body with the sanitized body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(sanitizedBody))
		}

		c.Next()
	}
}
