package mockdata

import (
	"net/http"
	"net/http/httptest"

	"github.com/porrporporrpor/covid-summary/model"
	"github.com/porrporporrpor/covid-summary/service"

	"github.com/stretchr/testify/mock"
)

type MockCovidCaseService struct{ mock.Mock }

func (m MockCovidCaseService) GetCovidCase(client service.HttpClientInterface) ([]model.CovidData, error) {
	args := m.Called(client)
	return args.Get(0).([]model.CovidData), args.Error(1)
}

type MockSummaryService struct{ mock.Mock }

func (m MockSummaryService) FindCountProvinceAndAgeGroup(data []model.CovidData) model.SummaryResponse {
	args := m.Called(data)
	return args.Get(0).(model.SummaryResponse)
}

type MockHttpClient struct {
	StatusCode int
	Response   string
}

func (c MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(c.StatusCode)
		_, _ = w.Write([]byte(c.Response))
	}))
	defer server.Close()
	return http.Get(server.URL)
}
