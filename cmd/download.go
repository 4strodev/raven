/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/4strodev/owl/git"
	"github.com/4strodev/raven/internals"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <REMOTE TEMPLATE>",
	Short: "Download a remote template",
	Long: `Every time you use a remote template raven saves it in a tmp folder.
To avoid download a remote template that you use it frequently it is a better option download it.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]
		var destination string

		regexParser := regexp.MustCompile("https?://")
		parsedRepo := regexParser.ReplaceAllString(repo, "")

		if len(args) == 2 {
			destination = args[0]
		}
		destination = path.Join(os.Getenv("HOME"), internals.LOCAL_TEMPLATES_DIRECTORY, parsedRepo)

		// check if template already was downloaded
		if Verbose {
			fmt.Printf("Downloading %s into %s\n", repo, destination)
		}
		err := git.Clone(repo, destination)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
