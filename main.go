package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-list",
	Short: "List files in a target folder",
	Run: func(cmd *cobra.Command, args []string) {
		targetFolder, _ := cmd.Flags().GetString("folder")
		outputFile, _ := cmd.Flags().GetString("output")

		files := listFiles(targetFolder)

		if outputFile != "" {
			saveToFile(outputFile, files)
		} else {
			printFiles(files)
		}
	},
}

func listFiles(folder string) []string {
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

func printFiles(files []string) {
	for _, file := range files {
		fmt.Println(file)
	}
}

func saveToFile(outputFile string, files []string) {
	targetFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer targetFile.Close()

	for _, file := range files {
		fmt.Fprintln(targetFile, file)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "::", "Result saved to", outputFile)
}

func init() {
	rootCmd.Flags().StringP("folder", "f", ".", "Target folder path")
	rootCmd.Flags().StringP("output", "o", "", "Output file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
