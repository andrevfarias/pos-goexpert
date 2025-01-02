package presenter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andrevfarias/go-expert/challenge5-stress-tester/internal/usecase"
)

type ReportPresenter struct {
}

func NewReportPresenter() *ReportPresenter {
	return &ReportPresenter{}
}

func (p *ReportPresenter) ToString(data usecase.GenerateReportUsecaseOutput) string {
	formattedDuration := data.TestDuration.Truncate(time.Millisecond).String()

	report := fmt.Sprintf("Total requests: %d\n", data.TotalRequests)
	report += fmt.Sprintf("Failed requests: %d\n", data.FailedRequests)
	report += fmt.Sprintf("Test duration: %s\n", formattedDuration)
	status200 := data.StatusCodeDistribution[200]
	report += fmt.Sprintf("\nSuccessful requests (HTTP200): %d (%.2f%%)\n\n", status200.Count, status200.Percentage)

	report += "Other status codes: "

	if len(data.StatusCodeDistribution) == 0 {
		return report + "Not available\n"
	}

	for code, item := range data.StatusCodeDistribution {
		if code != 200 {
			report += fmt.Sprintf("\n  HTTP%d:%3d (%5.2f%%)", code, item.Count, item.Percentage)
		}
	}

	return report + "\n"
}

func (p *ReportPresenter) ToJSON(data usecase.GenerateReportUsecaseOutput) string {
	jsonReport, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return "Error generating JSON report"
	}

	return string(jsonReport)
}
