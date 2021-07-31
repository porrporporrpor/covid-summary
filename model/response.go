package model

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
	PanicStatus   = "panic"
)

type SummaryResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

type HttpResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
