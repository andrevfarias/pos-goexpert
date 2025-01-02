package cli

import (
	"os"

	"github.com/andrevfarias/go-expert/challenge5-stress-tester/internal/infra/cli/command"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stress-tester",
	Short: "It is a CLI tool for performing HTTP stress testing on web applications.",
	Long:  `This tool helps in generating concurrent http requests to a web server and generating a report with the test results.`,
	Run:   command.StressTestCmd,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "", "URL of the web server to be tested")
	rootCmd.Flags().Int("requests", 1, "Number of requests to be made")
	rootCmd.Flags().Int("concurrency", 1, "Number of concurrent requests to be made")
	rootCmd.Flags().String("output", "text", "Output format (text or json)")
}
