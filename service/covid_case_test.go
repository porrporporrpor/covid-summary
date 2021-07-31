package service_test

import (
	"errors"
	"fmt"
	"github.com/porrporporrpor/covid-summary/mockdata"
	"net/http"
	"testing"

	"github.com/porrporporrpor/covid-summary/model"
	"github.com/porrporporrpor/covid-summary/service"

	"github.com/openlyinc/pointy"
)

func TestGetCovidCase(t *testing.T) {
	t.Run("it should return covid case slice", func(t *testing.T) {
		expectedResponse := []model.CovidData{
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
		expectedError := error(nil)

		mockResponse := `{
	 "Data": [
	{
		  "ConfirmDate": "2021-05-04",
		  "No": null,
		  "Age": 51,
		  "Gender": "หญิง",
		  "GenderEn": "Female",
		  "Nation": null,
		  "NationEn": "China",
		  "Province": "Phrae",
		  "ProvinceId": 46,
		  "District": null,
		  "ProvinceEn": "Phrae",
		  "StatQuarantine": 5
	},
	{
		  "ConfirmDate": "2021-05-01",
		  "No": null,
		  "Age": 25,
		  "Gender": null,
		  "GenderEn": null,
		  "Nation": null,
		  "NationEn": "India",
		  "Province": "Phrae",
		  "ProvinceId": 46,
		  "District": null,
		  "ProvinceEn": "Phrae",
		  "StatQuarantine": 15
    },
	{
		  "ConfirmDate": "2021-05-01",
		  "No": null,
		  "Age": null,
		  "Gender": "หญิง",
		  "GenderEn": "Female",
		  "Nation": null,
		  "NationEn": null,
		  "Province": "Samut Songkhram",
		  "ProvinceId": 58,
		  "District": null,
		  "ProvinceEn": "Samut Songkhram",
		  "StatQuarantine": 11
		},
	{
		  "ConfirmDate": "2021-05-02",
		  "No": null,
		  "Age": 39,
		  "Gender": null,
		  "GenderEn": null,
		  "Nation": null,
		  "NationEn": "USA",
		  "Province": null,
		  "ProvinceId": null,
		  "District": null,
		  "ProvinceEn": null,
		  "StatQuarantine": 10
	},
	{
		  "ConfirmDate": null,
		  "No": null,
		  "Age": 86,
		  "Gender": "หญิง",
		  "GenderEn": "Female",
		  "Nation": null,
		  "NationEn": "Thailand",
		  "Province": "Chonburi",
		  "ProvinceId": 11,
		  "District": null,
		  "ProvinceEn": "Chonburi",
		  "StatQuarantine": null
	}
	 ]
	}`
		client := mockdata.MockHttpClient{
			StatusCode: http.StatusOK,
			Response:   mockResponse,
		}
		actualResponse, actualError := service.CovidCaseService{}.GetCovidCase(client)

		if expectedError != actualError {
			t.Error(fmt.Sprintf("expected error %v but got %v", expectedError, actualError))
		}
		if len(expectedResponse) != len(actualResponse) {
			t.Error(fmt.Sprintf("expected length %v but got %v", len(expectedResponse), len(actualResponse)))
		}
	})
	t.Run("it should return error cannot unmarshal request", func(t *testing.T) {
		var expectedResponse []model.CovidData
		expectedError := errors.New("invalid character 'i' looking for beginning of value")

		mockResponse := `invalid response`
		client := mockdata.MockHttpClient{
			StatusCode: http.StatusOK,
			Response:   mockResponse,
		}
		actualResponse, actualError := service.CovidCaseService{}.GetCovidCase(client)

		if expectedError.Error() != actualError.Error() {
			t.Error(fmt.Sprintf("expected error %v but got %v", expectedError, actualError))
		}
		if len(expectedResponse) != len(actualResponse) {
			t.Error(fmt.Sprintf("expected length %v but got %v", len(expectedResponse), len(actualResponse)))
		}
	})
}
