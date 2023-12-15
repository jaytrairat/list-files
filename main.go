package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
		isSplit, _ := cmd.Flags().GetBool("split")
		splitPosition, _ := cmd.Flags().GetInt("position")

		listFiles(targetFolder, regexPattern, isSplit, splitPosition, outputFile)
	},
}

func listFiles(folder string, regexPattern string, isSplit bool, splitPosition int, outputFile string) {
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var fileNames []string
	regex := regexp.MustCompile(regexPattern)
	for _, file := range files {
		targetFileName := file.Name()
		fileName := targetFileName
		if regex.MatchString(targetFileName) {
			if isSplit {
				splitedText := strings.Split(targetFileName, " ")

				targetSplitPosition := splitPosition
				if splitPosition >= len(splitedText) {
					targetSplitPosition = len(splitedText) - 1
				}
				fileName = splitedText[targetSplitPosition]
			}

			fileNames = append(fileNames, fileName)
		}
	}
	if outputFile != "" {
		targetFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer targetFile.Close()
		for _, fileName := range fileNames {
			fmt.Fprintln(targetFile, fileName)
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "::", "Result saved to", outputFile)
	}

}

func init() {
	rootCmd.Flags().StringP("folder", "f", ".", "Target folder path")
	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.Flags().StringP("regex", "r", "", "Pattern of file name")
	rootCmd.Flags().BoolP("split", "s", false, "Is split base on regex")
	rootCmd.Flags().IntP("position", "p", 0, "Position to be splited")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
