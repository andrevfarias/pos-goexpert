package usecase

import "time"

type GenerateReportUsecaseInput struct {
	Requests       int
	FailedRequests int
	StatusCode     []int
	TestDuration   time.Duration
}

type StatusCodeDistributionItem struct {
	Count      int     `json:"count"`
	Percentage float64 `json:"percentage"`
}

type GenerateReportUsecaseOutput struct {
	TotalRequests          int                                `json:"total_requests"`
	FailedRequests         int                                `json:"failed_requests"`
	TestDuration           time.Duration                      `json:"test_duration_ms"`
	StatusCodeDistribution map[int]StatusCodeDistributionItem `json:"status_code_distribution"`
}

type GenerateReportUsecaseInterface interface {
	Execute(input GenerateReportUsecaseInput) (GenerateReportUsecaseOutput, error)
}

type GenerateReportUsecase struct{}

func NewGenerateReportUsecase() GenerateReportUsecaseInterface {
	return &GenerateReportUsecase{}
}

func (u *GenerateReportUsecase) Execute(input GenerateReportUsecaseInput) (GenerateReportUsecaseOutput, error) {
	output := GenerateReportUsecaseOutput{
		TotalRequests:          input.Requests,
		FailedRequests:         input.FailedRequests,
		TestDuration:           input.TestDuration,
		StatusCodeDistribution: make(map[int]StatusCodeDistributionItem),
	}

	for _, code := range input.StatusCode {
		count := output.StatusCodeDistribution[code].Count + 1
		output.StatusCodeDistribution[code] = StatusCodeDistributionItem{
			Count:      count,
			Percentage: (float64(count) / float64(output.TotalRequests)) * 100,
		}
	}

	return output, nil
}
