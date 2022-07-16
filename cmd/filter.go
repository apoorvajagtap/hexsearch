/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/fatih/color"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "List the files or sentences that contains specified string",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range os.Args {
			if c == "--list" {
				fetchFile(args[0])
			}

			if c == "--print" {
				displayChat(args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filterCmd.PersistentFlags().String("foo", "", "A help for foo")
	var list, print bool
	filterCmd.Flags().BoolVarP(&list, "list", "l", false, "List the chat files with specified string")
	filterCmd.Flags().BoolVarP(&print, "print", "p", false, "Print out the complete messages on screen")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchFile(keyword string) {
	if err := os.Chdir(homeDir + dirLocation); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n", "Here's is the list of files with relevant keyword")
	// color.Cyan("Here's is the list of files with relevant keyword")
	files, _ := os.ReadDir(homeDir + dirLocation)
	for _, ele := range files {
		f, err := os.Open(ele.Name())
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), keyword) {
				fmt.Println(ele.Name())
				break
			}
		}
	}
}

func displayChat(keyword string) {
	if err := os.Chdir(homeDir + dirLocation); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Here's are the sentences with relevant keyword")
	files, _ := os.ReadDir(homeDir + dirLocation)
	for _, ele := range files {
		f, err := os.Open(ele.Name())
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {

			if strings.Contains(scanner.Text(), keyword) {
				fmt.Println(scanner.Text())
			}

		}
	}
}
