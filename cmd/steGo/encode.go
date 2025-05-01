package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thule/steGo/pkg/stego"
)

var (
	inputFile   string
	outputFile  string
	message     string
	messageFile string
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a message into an image",
	Long: `Encode a message into an image using steganography.
The message can be provided directly or from a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" || outputFile == "" {
			fmt.Println("Error: input and output files are required")
			os.Exit(1)
		}

		if message == "" && messageFile == "" {
			fmt.Println("Error: either message or message file must be provided")
			os.Exit(1)
		}

		var msg string

		if messageFile != "" {
			data, err := os.ReadFile(messageFile)
			if err != nil {
				fmt.Printf("Error reading message file: %v\n", err)
				os.Exit(1)
			}
			msg = string(data)
		} else {
			msg = message
		}

		if err := stego.EncodeMessage(inputFile, outputFile, msg); err != nil {
			fmt.Printf("Error encoding message: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Message successfully encoded into %s\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input image file (required)")
	encodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output image file (required)")
	encodeCmd.Flags().StringVarP(&message, "message", "m", "", "message to encode")
	encodeCmd.Flags().StringVarP(&messageFile, "file", "f", "", "file containing message to encode")

	encodeCmd.MarkFlagRequired("input")
	encodeCmd.MarkFlagRequired("output")
}
