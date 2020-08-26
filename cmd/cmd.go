package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gvmCmd = &cobra.Command{
	Use:   "gvm",
	Short: "GVM is a java virtual machine written in Golang",
}

func Execute() error {
	return gvmCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		viper.AutomaticEnv()
	})
}
