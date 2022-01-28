package main

import (
	"log"

	"github.com/brockweaver/hxu/cmd"
)

var buildVersion string

func GetBuildInfo() string {
	return buildVersion
}

func main() {
	rootCmd := cmd.CobraCommands(buildVersion)
	err2 := rootCmd.Execute()
	if err2 != nil {
		log.Fatalf(err2.Error())
	}
}
