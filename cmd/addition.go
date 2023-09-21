/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// additionCmd represents the addition command
var additionCmd = &cobra.Command{
	Use:   "addition",
	Short: "perform addition of two numbers",
	Long: `Example usage:
  	golang-cli addition 1 2
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid number:", args[0])
			os.Exit(1)
		}

		num2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid number:", args[1])
			os.Exit(1)
		}

		result := num1 + num2
		fmt.Printf("Result: %d + %d = %d\n", num1, num2, result)
	},
}

func init() {
	rootCmd.AddCommand(additionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// additionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// additionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
