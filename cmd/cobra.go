package cmd

import (
	"fmt"

	"github.com/brockweaver/hxu/cmd/file"
	"github.com/spf13/cobra"
)

var buildInfo string

func CobraCommands(buildVersion string) *cobra.Command {

	buildInfo = buildVersion

	rootCmd := &cobra.Command{
		Use: "hxu",
	}

	file.CobraCommands(rootCmd)

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "reports the version of hxu",
		Run:   cobraVersion,
	}
	rootCmd.AddCommand(versionCmd)

	return rootCmd

}

func cobraVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("version: %v\n", buildInfo)
}
