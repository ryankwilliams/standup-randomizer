package cmd

import (
	"os"
	"standup-randomizer/pkg/randomizer"

	"github.com/spf13/cobra"
)

// TODO: Accept positional argument to get file with list of team members

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "standup-randomizer",
	Short: "Generates daily scrum team member stand up order",
	Long:  `Generates daily scrum team member stand up order`,
	Run: func(cmd *cobra.Command, args []string) {
		randomizer.GenerateOrder()
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.standup-randomizer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
