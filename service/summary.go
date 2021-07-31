package service

import (
	"github.com/porrporporrpor/covid-summary/model"
)

type SummaryServiceInterface interface {
	FindCountProvinceAndAgeGroup(data []model.CovidData) model.SummaryResponse
}

type SummaryService struct{}

func (s SummaryService) FindCountProvinceAndAgeGroup(data []model.CovidData) model.SummaryResponse {
	province := make(map[string]int)
	ageGroup := make(map[string]int)

	for _, d := range data {
		provinceData := d.Province
		ageData := d.Age

		if provinceData != nil {
			province[*provinceData] = province[*provinceData] + 1
		}

		if ageData != nil {
			ageRange := CheckAgeGroup(*ageData)
			ageGroup[ageRange] = ageGroup[ageRange] + 1
		} else {
			ageGroup["N/A"] = ageGroup["N/A"] + 1
		}
	}
	return model.SummaryResponse{Province: province, AgeGroup: ageGroup}
}

func CheckAgeGroup(age int) string {
	if age > 60 {
		return "61+"
	}
	if age >= 31 && age <= 60 {
		return "31-60"
	}
	if age >= 0 && age <= 30 {
		return "0-30"
	}
	return "N/A"
}
