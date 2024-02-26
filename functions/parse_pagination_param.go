package functions

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePaginationParams(c *gin.Context) (int64, int64, error) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	var offset int64 = 0 // Default offset value
	var err error
	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid offset format: %w", err)
		}
	}

	var limit int64 = 10 // Default limit value
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid limit format: %w", err)
		}
	}

	return offset, limit, nil
}
