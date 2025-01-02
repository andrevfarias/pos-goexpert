package command

import (
	"github.com/andrevfarias/go-expert/challenge5-stress-tester/internal/infra/cli/presenter"
	"github.com/andrevfarias/go-expert/challenge5-stress-tester/internal/usecase"
	"github.com/spf13/cobra"
)

func StressTestCmd(cmd *cobra.Command, args []string) {
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		cmd.Println("Error getting URL flag")
	}
	if url == "" {
		cmd.Println("URL flag is required")
		return
	}

	requests, err := cmd.Flags().GetInt("requests")
	if err != nil {
		cmd.Println("Error getting requests flag")
	}
	if requests <= 0 {
		cmd.Println("Requests flag should be greater than 0")
		return
	}

	concurrency, err := cmd.Flags().GetInt("concurrency")
	if err != nil {
		cmd.Println("Error getting concurrency flag")
	}
	if concurrency <= 0 {
		cmd.Println("Concurrency flag should be greater than 0")
		return
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Println("Error getting output flag")
	}
	if outputFormat != "text" && outputFormat != "json" {
		cmd.Println("Output flag should be text or json")
		return
	}

	performStressTestUsecaseInput := usecase.PerformStressTestUsecaseInput{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
	performStressTestUsecase := usecase.NewPerformStressTestUsecase()
	performStressTestUsecaseOutput, err := performStressTestUsecase.Execute(performStressTestUsecaseInput)

	if err != nil {
		cmd.Println("Error executing stress test ->", err)
		return
	}

	generateReportUsecaseInput := usecase.GenerateReportUsecaseInput(performStressTestUsecaseOutput)
	generateReportUsecase := usecase.NewGenerateReportUsecase()
	generateReportUsecaseOutput, err := generateReportUsecase.Execute(generateReportUsecaseInput)

	if err != nil {
		cmd.Println("Error generating report ->", err)
		return
	}

	reportPresenter := presenter.NewReportPresenter()
	switch outputFormat {
	case "json":
		cmd.Println(reportPresenter.ToJSON(generateReportUsecaseOutput))
	default:
		cmd.Println(reportPresenter.ToString(generateReportUsecaseOutput))
	}
}
