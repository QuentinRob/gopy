package cmd

import (
	"github.com/spf13/cobra"
)

const RECURSIVE_FLAG = "recursive"

var (
	Recursive bool
	rootCmd   = &cobra.Command{
		Use:   "gp",
		Short: "Golang implementation of the copy command with more feedback about the status",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(2)),
		PreRunE: Validate,
		RunE:    Copy,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Recursive, RECURSIVE_FLAG, "r", false, "Recursive flag to copy folders instead of files")
}
