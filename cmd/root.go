package cmd

import (
	"fmt"
	"github.com/adamkasztenny/slack-repeat-bot/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "slack-repeat-bot",
	Version: "1.0.0",
	Short:   "A Slack bot that repeats a word multiple times",
	Run: func(cmd *cobra.Command, args []string) {
		validateArgs(args)
		initializeLogger()

		token := args[0]
		repeatBotAPI := new(api.API)
		repeatBotAPI.Initialize(token)
		repeatBotAPI.Listen()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateArgs(args []string) {
	if len(args) < 1 {
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: slack-repeat-bot [api token]")
	os.Exit(1)
}

func initializeLogger() {
	formatter := new(log.TextFormatter)
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
}
