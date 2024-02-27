package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "order-packs-calculator",
		Short: "order-packs-calculator for calculating the least wasted packages for order items based on predefined package sizes configuration",
		Long:  "order-packs-calculator for calculating the least wasted packages for order items based on predefined package sizes configuration",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// Execute ...
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
