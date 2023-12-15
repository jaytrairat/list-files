package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-list",
	Short: "List files in a target folder",
	Run: func(cmd *cobra.Command, args []string) {
		targetFolder, _ := cmd.Flags().GetString("folder")
		outputFile, _ := cmd.Flags().GetString("output")
		regexPattern, _ := cmd.Flags().GetString("regex")

		files := listFiles(targetFolder, regexPattern)

		if outputFile != "" {
			saveToFile(outputFile, files, regexPattern)
		} else {
			printFiles(files)
		}
	},
}

func listFiles(folder string, regexPattern string) []string {
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var fileNames []string
	regex := regexp.MustCompile(regexPattern)
	for _, file := range files {
		if regex.MatchString(file.Name()) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

func printFiles(files []string) {
	for _, file := range files {
		fmt.Println(file)
	}
}

func saveToFile(outputFile string, files []string, regexPattern string) {
	targetFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer targetFile.Close()

	regex := regexp.MustCompile(regexPattern)
	for _, file := range files {
		if regex.MatchString(file) {
			fmt.Fprintln(targetFile, file)
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "::", "Result saved to", outputFile)
}

func init() {
	rootCmd.Flags().StringP("folder", "f", ".", "Target folder path")
	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.Flags().StringP("regex", "r", "", "Pattern of file name")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
