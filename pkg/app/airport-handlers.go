package app

import (
	"golang-couchbase/pkg/api"
	"log"

	"github.com/gin-gonic/gin"
)

func (s *Server) SearchAirport() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchKey := c.Query("search")

		if searchKey == "" {
			log.Printf("handler error: %s", api.ErrorAirportSearchCriteriaRequired)
			writeJsonFailure(c, api.ErrorAirportSearchCriteriaRequired.Error(), nil, nil)
			return
		}

		respData, err := s.airportService.Search(searchKey)

		if err != nil {
			log.Printf("handler error: %s", err)
			writeJsonFailure(c, err.Error(), nil, nil)
			return
		}

		writeResponseSuccess(c, respData)
	}
}
