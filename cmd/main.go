package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/khenidak/ip-thing/pkg/options"
)

// set by build flags
var Version string
var BuildTime string

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	opts := options.MakeOptions(Version, BuildTime)
	theApp := newApp(opts)
	theApp.Execute()
}

func newApp(opts *options.Options) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "ip-thing",
		Short: "does the ip thing",
		Long:  `does the ip thing`,
	}

	cmd.AddCommand(newVersionCommand(opts))
	cmd.AddCommand(newRunCommand(opts))

	return cmd
}
