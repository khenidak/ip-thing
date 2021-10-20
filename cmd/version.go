package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/khenidak/ip-thing/pkg/options"
)

func newVersionCommand(opts *options.Options) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of k8s api generator",
		Long:  `Print the version number of k8s api generator`,

		Run: func(cmd *cobra.Command, args []string) {
			printVersionInfo(opts)
		},
	}
}

func printVersionInfo(opts *options.Options) {
	fmt.Printf("Version %s at %s\n", opts.Version, opts.BuildTime)
}
