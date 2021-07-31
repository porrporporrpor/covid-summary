package service

import (
	"crypto/tls"
	"encoding/json"
	"github.com/porrporporrpor/covid-summary/model"
	"io"
	"net/http"
)

type CovidCaseServiceInterface interface {
	GetCovidCase(client HttpClientInterface) ([]model.CovidData, error)
}

type CovidCaseService struct {
}

type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c CovidCaseService) GetCovidCase(client HttpClientInterface) ([]model.CovidData, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := "https://static.wongnai.com/devinterview/covid-cases.json"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var covidCase model.CovidRequestBody
	err = json.Unmarshal(body, &covidCase)
	if err != nil {
		return nil, err
	}

	return covidCase.Data, nil
}
