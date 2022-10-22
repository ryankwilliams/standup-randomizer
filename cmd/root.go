package cmd

import (
	"os"
	"standup-randomizer/pkg/randomizer"

	"github.com/spf13/cobra"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "standup-randomizer",
	Short: "Generates daily scrum team member stand up order",
	Long:  `Generates daily scrum team member stand up order`,
	Run: func(cmd *cobra.Command, args []string) {
		randomizer.GenerateOrder(configFile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "",
		"config file (default is $HOME/.standup-randomizer.yaml)")
}
