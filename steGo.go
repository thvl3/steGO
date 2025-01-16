package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "strings"
)

// Main function to handle CLI arguments
func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage:")
        fmt.Println("  Encode: steGo encode input.png output.png \"Hidden message\"")
        fmt.Println("  Encode with file: steGo encode input.png output.png -file secret.txt")
        fmt.Println("  Decode: steGo decode output.png")
        fmt.Println("  Decode to file: steGo decode output.png -file extracted.txt")
        return
    }

    mode := os.Args[1] // "encode" or "decode"
    input := os.Args[2] // Input PNG
    output := ""
    message := ""

    if mode == "encode" {
        if len(os.Args) < 4 {
            fmt.Println("Error: Missing output file.")
            return
        }
        output = os.Args[3] // Output PNG

        if len(os.Args) > 4 && os.Args[4] == "-file" {
            if len(os.Args) < 6 {
                fmt.Println("Error: Missing input file for encoding.")
                return
            }
            filePath := os.Args[5] // File to read data from
            data, err := os.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err)
                return
            }
            message = string(data)
        } else {
            message = os.Args[4] // Raw text message
        }

        encodeMessage(input, output, message)
    } else if mode == "decode" {
        extracted := decodeMessage(input)

        if len(os.Args) > 3 && os.Args[3] == "-file" {
            if len(os.Args) < 5 {
                fmt.Println("Error: Missing output file for extracted text.")
                return
            }
            filePath := os.Args[4]
            err := os.WriteFile(filePath, []byte(extracted), 0644)
            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
            fmt.Println("Hidden message extracted to:", filePath)
        } else {
            fmt.Println("Hidden message:", extracted)
        }
    }
}

// Encode a hidden message inside a PNG image
func encodeMessage(inputFile, outputFile, message string) {
    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening image:", err)
        return
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error decoding image:", err)
        return
    }

    bounds := img.Bounds()
    imgWidth := bounds.Dx()
    imgHeight := bounds.Dy()
    totalPixels := imgWidth * imgHeight

    // Convert message to binary
    messageBytes := []byte(message)
    messageBytes = append(messageBytes, 0) // Null terminator to signal end of message
    messageLength := len(messageBytes)

    if messageLength*8 > totalPixels {
        fmt.Println("Error: Message is too large to fit in the image!")
        return
    }

    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outFile.Close()

    png.Encode(outFile, hideMessage(img, messageBytes))
    fmt.Println("Message successfully embedded in", outputFile)
}

// Extract a hidden message from a PNG image
func decodeMessage(inputFile string) string {
    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening image:", err)
        return ""
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error decoding image:", err)
        return ""
    }

    extractedBytes := extractHiddenMessage(img)
    return string(extractedBytes)
}

// Hide a message inside an image using LSB
func hideMessage(img image.Image, messageBytes []byte) *image.RGBA {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)

    msgIndex := 0
    bitIndex := 0

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            originalColor := img.At(x, y)
            r, g, b, a := originalColor.RGBA()
            r, g, b = r>>8, g>>8, b>>8 // Convert to 8-bit values

            // Modify LSB of red channel to store message bits
            if msgIndex < len(messageBytes) {
                bit := (messageBytes[msgIndex] >> (7 - bitIndex)) & 1
                r = (r & 0xFE) | bit

                bitIndex++
                if bitIndex == 8 {
                    bitIndex = 0
                    msgIndex++
                }
            }

            newImg.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a >> 8)})
        }
    }
    return newImg
}

// Extract a hidden message from an image using LSB
func extractHiddenMessage(img image.Image) []byte {
    bounds := img.Bounds()
    var extractedBits []byte
    var currentByte byte = 0
    bitIndex := 0

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, _, _, _ := img.At(x, y).RGBA()
            r = r >> 8 // Convert to 8-bit value

            // Extract LSB from red channel
            currentByte = (currentByte << 1) | byte(r&1)
            bitIndex++

            if bitIndex == 8 {
                if currentByte == 0 { // Null terminator signals end of message
                    return extractedBits
                }
                extractedBits = append(extractedBits, currentByte)
                bitIndex = 0
                currentByte = 0
            }
        }
    }
    return extractedBits
}
