package app

import (
	"golang-couchbase/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		writeJsonSuccess(c, "weight tracker API running smoothly", nil, nil)
	}
}

func writeJsonResponse(c *gin.Context, code int, response interface{}) {
	c.Header("Content-Type", "application/json")

	c.JSON(code, response)
}

func writeResponseSuccess(c *gin.Context, response interface{}) {
	writeJsonResponse(c, http.StatusOK, response)
}

func writeResponseFailure(c *gin.Context, response interface{}) {
	writeJsonResponse(c, http.StatusInternalServerError, response)
}

func writeJsonResponseWithContext(c *gin.Context, code int, status string, data interface{}, context api.Context) {
	c.Header("Content-Type", "application/json")

	response := map[string]interface{}{
		"status":  status,
		"data":    data,
		"context": context,
	}

	c.JSON(code, response)
}

func writeJsonSuccess(c *gin.Context, status string, data interface{}, context api.Context) {
	writeJsonResponseWithContext(c, http.StatusOK, status, data, context)
}

func writeJsonFailure(c *gin.Context, status string, data interface{}, context api.Context) {
	writeJsonResponseWithContext(c, http.StatusInternalServerError, status, data, context)
}
