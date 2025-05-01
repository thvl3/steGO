package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thule/steGo/pkg/stego"
)

var (
	decodeInputFile  string
	decodeOutputFile string
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a message from an image",
	Long: `Decode a message from an image using steganography.
The message can be output to stdout or saved to a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if decodeInputFile == "" {
			fmt.Println("Error: input file is required")
			os.Exit(1)
		}

		msg, err := stego.DecodeMessage(decodeInputFile)
		if err != nil {
			fmt.Printf("Error decoding message: %v\n", err)
			os.Exit(1)
		}

		if decodeOutputFile != "" {
			if err := os.WriteFile(decodeOutputFile, []byte(msg), 0644); err != nil {
				fmt.Printf("Error writing to output file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Message successfully decoded to %s\n", decodeOutputFile)
		} else {
			fmt.Println("Decoded message:")
			fmt.Println(msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringVarP(&decodeInputFile, "input", "i", "", "input image file (required)")
	decodeCmd.Flags().StringVarP(&decodeOutputFile, "output", "o", "", "output file (optional)")

	decodeCmd.MarkFlagRequired("input")
}
