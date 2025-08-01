package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple CLI archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
