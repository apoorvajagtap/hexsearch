/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hexsearch",
	Short: "A CLI tool to fetch the relevant hexchat messages",
	Long: `This tool allows you to get the complete chat history of any specific user, 
or get specific strings fetched out from the hexchat's chats`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	homeDir, _ := os.UserHomeDir()
	// 	dirLocation := homeDir + "/.config/hexchat/scrollback/red hat"

	// 	// fmt.Print(dirLocation)
	// },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hexsearch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// var chat string
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().String(chat, "", "Name of the user, you wish to determine the chat for.")
	// userName := flag.String("chat", "", "Name of the user, you wish to find the chat for.")
	// grepDetail := flag.String("filter", "", "Specify the string, you wish to search among all the chats.")
	// flag.Parse()

	// if *userName != "" {
	// 	chatOptions.GetChat(*userName)
	// }

	// // if *grepDetail != "" {
	// chatOptions.FetchSent(*grepDetail)
	// }

	// fmt.Println(*userName, *grepDetail)
}
