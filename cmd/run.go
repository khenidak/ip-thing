package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/khenidak/ip-thing/pkg/options"
)

func newRunCommand(opts *options.Options) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "runs the ip-thing",
		Long:  `runs the ip thing`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("run..")
		},
	}
}
