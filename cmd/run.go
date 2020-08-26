package cmd

import (
	"github.com/spf13/cobra"
	"gvm/machine"
)

var command = &machine.Command{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run java class",
	Run: func(cmd *cobra.Command, args []string) {
		command.Args = args
		machine.Run(command)
	},
}

func init() {
	runCmd.Flags().StringVarP(&command.JarFile, "jar", "", "", "path to jar file")
	gvmCmd.AddCommand(runCmd)
}
