package service_test

import (
	"fmt"
	"github.com/openlyinc/pointy"
	"github.com/porrporporrpor/covid-summary/model"
	"github.com/porrporporrpor/covid-summary/service"
	"reflect"
	"testing"
)

func TestCheckAgeGroup(t *testing.T) {
	t.Run("it should return 60+", func(t *testing.T) {
		expectedResponse := "60+"
		actualResponse := service.CheckAgeGroup(90)
		if expectedResponse != actualResponse {
			t.Error(fmt.Sprintf("expected %v but got %v", expectedResponse, actualResponse))
		}
	})

	t.Run("it should return 31-60", func(t *testing.T) {
		expectedResponse := "31-60"
		actualResponse := service.CheckAgeGroup(60)
		if expectedResponse != actualResponse {
			t.Error(fmt.Sprintf("expected %v but got %v", expectedResponse, actualResponse))
		}
	})

	t.Run("it should return 0-30", func(t *testing.T) {
		expectedResponse := "0-30"
		actualResponse := service.CheckAgeGroup(30)
		if expectedResponse != actualResponse {
			t.Error(fmt.Sprintf("expected %v but got %v", expectedResponse, actualResponse))
		}
	})

	t.Run("it should return N/A", func(t *testing.T) {
		expectedResponse := "N/A"
		actualResponse := service.CheckAgeGroup(-1)
		if expectedResponse != actualResponse {
			t.Error(fmt.Sprintf("expected %v but got %v", expectedResponse, actualResponse))
		}
	})
}

func TestFindCountProvinceAndAgeGroup(t *testing.T) {
	inputData := []model.CovidData{
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

	t.Run("it should return stat count province and age group", func(t *testing.T) {
		expectedResponse := model.SummaryResponse{
			Province: map[string]int{
				"Phrae":           2,
				"Samut Songkhram": 1,
				"Chonburi":        1,
			},
			AgeGroup: map[string]int{
				"0-30":  1,
				"31-60": 2,
				"60+":   1,
				"N/A":   1,
			},
		}
		actualResponse := service.SummaryService{}.FindCountProvinceAndAgeGroup(inputData)
		if !reflect.DeepEqual(expectedResponse, actualResponse) {
			t.Error(fmt.Sprintf("expected %v but got %v", expectedResponse, actualResponse))
		}
	})
}
