package api_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/porrporporrpor/covid-summary/api"
	"github.com/porrporporrpor/covid-summary/mockdata"
	"github.com/porrporporrpor/covid-summary/model"

	"github.com/gin-gonic/gin"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestSummaryAPI_SummaryAPI(t *testing.T) {
	t.Run("it should return stat data with status OK", func(t *testing.T) {
		expectedStatusCode := http.StatusOK
		expectedResponse := `{"status":"success","data":{"Province":{"Chonburi":1,"Phrae":2,"Samut Songkhram":1},"AgeGroup":{"0-30":1,"31-60":2,"61+":1,"N/A":1}}}`

		covidCaseService := mockdata.MockCovidCaseService{}
		covidCase := []model.CovidData{
			{
				ConfirmDate:    pointy.String("2021-05-04"),
				No:             nil,
				Age:            pointy.Int(51),
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       pointy.String("China"),
				Province:       pointy.String("Phrae"),
				ProvinceId:     pointy.Int(46),
				District:       nil,
				ProvinceEn:     pointy.String("Phrae"),
				StatQuarantine: pointy.Int(5),
			},
			{
				ConfirmDate:    pointy.String("2021-05-01"),
				No:             nil,
				Age:            pointy.Int(25),
				Gender:         nil,
				GenderEn:       nil,
				Nation:         nil,
				NationEn:       pointy.String("India"),
				Province:       pointy.String("Phrae"),
				ProvinceId:     pointy.Int(46),
				District:       nil,
				ProvinceEn:     pointy.String("Phrae"),
				StatQuarantine: pointy.Int(15),
			},
			{
				ConfirmDate:    pointy.String("2021-05-01"),
				No:             nil,
				Age:            nil,
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       nil,
				Province:       pointy.String("Samut Songkhram"),
				ProvinceId:     pointy.Int(58),
				District:       nil,
				ProvinceEn:     pointy.String("Samut Songkhram"),
				StatQuarantine: pointy.Int(11),
			},
			{
				ConfirmDate:    pointy.String("2021-05-02"),
				No:             nil,
				Age:            pointy.Int(39),
				Gender:         nil,
				GenderEn:       nil,
				Nation:         nil,
				NationEn:       pointy.String("USA"),
				Province:       nil,
				ProvinceId:     nil,
				District:       nil,
				ProvinceEn:     nil,
				StatQuarantine: pointy.Int(10),
			},
			{
				ConfirmDate:    nil,
				No:             nil,
				Age:            pointy.Int(86),
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       pointy.String("Thailand"),
				Province:       pointy.String("Chonburi"),
				ProvinceId:     pointy.Int(11),
				District:       nil,
				ProvinceEn:     pointy.String("Chonburi"),
				StatQuarantine: nil,
			},
		}
		client := http.Client{}
		covidCaseService.On("GetCovidCase", &client).Return(covidCase, nil)

		summaryStat := model.SummaryResponse{
			Province: map[string]int{
				"Phrae":           2,
				"Samut Songkhram": 1,
				"Chonburi":        1,
			},
			AgeGroup: map[string]int{
				"0-30":  1,
				"31-60": 2,
				"61+":   1,
				"N/A":   1,
			},
		}
		summaryService := mockdata.MockSummaryService{}
		summaryService.On("FindCountProvinceAndAgeGroup", covidCase).Return(summaryStat)

		summaryAPI := api.SummaryAPI{
			CovidCaseService: covidCaseService,
			SummaryService:   summaryService,
		}

		request := httptest.NewRequest(http.MethodGet, "/covid/summary", nil)
		request.Header.Add("Content-Type", "application/json")
		writer := httptest.NewRecorder()

		server := gin.New()
		server.GET("/covid/summary", summaryAPI.SummaryAPI)
		server.ServeHTTP(writer, request)
		response := writer.Result()

		actualStatusCode := response.StatusCode
		actualResponse, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, expectedStatusCode, actualStatusCode)
		assert.Equal(t, expectedResponse, string(actualResponse))
	})

	t.Run("it should return error with status InternalServerError", func(t *testing.T) {
		expectedStatusCode := http.StatusInternalServerError
		expectedResponse := `{"status":"fail","data":"cannot get covid data"}`

		covidCaseService := mockdata.MockCovidCaseService{}
		covidCase := []model.CovidData{
			{
				ConfirmDate:    pointy.String("2021-05-04"),
				No:             nil,
				Age:            pointy.Int(51),
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       pointy.String("China"),
				Province:       pointy.String("Phrae"),
				ProvinceId:     pointy.Int(46),
				District:       nil,
				ProvinceEn:     pointy.String("Phrae"),
				StatQuarantine: pointy.Int(5),
			},
			{
				ConfirmDate:    pointy.String("2021-05-01"),
				No:             nil,
				Age:            pointy.Int(25),
				Gender:         nil,
				GenderEn:       nil,
				Nation:         nil,
				NationEn:       pointy.String("India"),
				Province:       pointy.String("Phrae"),
				ProvinceId:     pointy.Int(46),
				District:       nil,
				ProvinceEn:     pointy.String("Phrae"),
				StatQuarantine: pointy.Int(15),
			},
			{
				ConfirmDate:    pointy.String("2021-05-01"),
				No:             nil,
				Age:            nil,
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       nil,
				Province:       pointy.String("Samut Songkhram"),
				ProvinceId:     pointy.Int(58),
				District:       nil,
				ProvinceEn:     pointy.String("Samut Songkhram"),
				StatQuarantine: pointy.Int(11),
			},
			{
				ConfirmDate:    pointy.String("2021-05-02"),
				No:             nil,
				Age:            pointy.Int(39),
				Gender:         nil,
				GenderEn:       nil,
				Nation:         nil,
				NationEn:       pointy.String("USA"),
				Province:       nil,
				ProvinceId:     nil,
				District:       nil,
				ProvinceEn:     nil,
				StatQuarantine: pointy.Int(10),
			},
			{
				ConfirmDate:    nil,
				No:             nil,
				Age:            pointy.Int(86),
				Gender:         pointy.String("หญิง"),
				GenderEn:       pointy.String("Female"),
				Nation:         nil,
				NationEn:       pointy.String("Thailand"),
				Province:       pointy.String("Chonburi"),
				ProvinceId:     pointy.Int(11),
				District:       nil,
				ProvinceEn:     pointy.String("Chonburi"),
				StatQuarantine: nil,
			},
		}
		client := http.Client{}
		covidCaseService.
			On("GetCovidCase", &client).
			Return([]model.CovidData{}, errors.New("cannot get covid data"))

		summaryStat := model.SummaryResponse{
			Province: map[string]int{
				"Phrae":           2,
				"Samut Songkhram": 1,
				"Chonburi":        1,
			},
			AgeGroup: map[string]int{
				"0-30":  1,
				"31-60": 2,
				"61+":   1,
				"N/A":   1,
			},
		}
		summaryService := mockdata.MockSummaryService{}
		summaryService.On("FindCountProvinceAndAgeGroup", covidCase).Return(summaryStat)

		summaryAPI := api.SummaryAPI{
			CovidCaseService: covidCaseService,
			SummaryService:   summaryService,
		}

		request := httptest.NewRequest(http.MethodGet, "/covid/summary", nil)
		request.Header.Add("Content-Type", "application/json")
		writer := httptest.NewRecorder()

		server := gin.New()
		server.GET("/covid/summary", summaryAPI.SummaryAPI)
		server.ServeHTTP(writer, request)
		response := writer.Result()

		actualStatusCode := response.StatusCode
		actualResponse, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, expectedStatusCode, actualStatusCode)
		assert.Equal(t, expectedResponse, string(actualResponse))
	})
}
