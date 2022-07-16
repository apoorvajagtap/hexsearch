/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Print Chat with specific user",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		getChat(args[0])
	},
}

var homeDir, _ = os.UserHomeDir()

const (
	dirLocation = "/.config/hexchat/scrollback/red hat/"
)

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getChat(user string) {
	fileName := homeDir + dirLocation + user + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
