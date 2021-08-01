package router

import (
	"github.com/porrporporrpor/covid-summary/api"
	"github.com/porrporporrpor/covid-summary/middleware"
	"github.com/porrporporrpor/covid-summary/service"

	"github.com/gin-gonic/gin"
)

func Boot(r *gin.Engine) error {
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(middleware.CustomRecovery))
	r.Use(middleware.CORS())

	covidGroupRoute(r.Group("/covid"))

	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}

func covidGroupRoute(r *gin.RouterGroup) {
	summaryAPI := api.SummaryAPI{
		SummaryService:   service.SummaryService{},
		CovidCaseService: service.CovidCaseService{},
	}

	r.GET("/summary", summaryAPI.SummaryAPI)
}
