package api

import (
	"net/http"

	"github.com/porrporporrpor/covid-summary/model"
	"github.com/porrporporrpor/covid-summary/service"

	"github.com/gin-gonic/gin"
)

type SummaryAPI struct {
	CovidCaseService service.CovidCaseServiceInterface
	SummaryService   service.SummaryServiceInterface
}

func (s *SummaryAPI) SummaryAPI(ctx *gin.Context) {
	client := http.Client{}
	covidCase, err := s.CovidCaseService.GetCovidCase(&client)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.HttpResponse{Status: model.FailStatus, Data: err.Error()})
		return
	}

	stat := s.SummaryService.FindCountProvinceAndAgeGroup(covidCase)

	ctx.JSON(http.StatusOK, model.HttpResponse{Status: model.SuccessStatus, Data: stat})
}
