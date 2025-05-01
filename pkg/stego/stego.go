package stego

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// EncodeMessage encodes a message into an image using LSB steganography
func EncodeMessage(inputFile, outputFile, message string) error {
	// Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		return NewError(ErrFileOperation, ErrMsgFileOperation, err)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return NewError(ErrImageProcessing, ErrMsgImageProcessing, err)
	}

	// Validate image format
	if _, ok := img.(*image.RGBA); !ok {
		return NewError(ErrInvalidImageFormat, ErrMsgInvalidImageFormat, nil)
	}

	// Validate message size
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	totalPixels := imgWidth * imgHeight

	messageBytes := []byte(message)
	messageBytes = append(messageBytes, 0) // Null terminator
	messageLength := len(messageBytes)

	if messageLength*8 > totalPixels {
		return NewError(ErrMessageTooLarge, ErrMsgMessageTooLarge, nil)
	}

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return NewError(ErrFileOperation, ErrMsgFileOperation, err)
	}
	defer outFile.Close()

	// Hide message
	newImg := hideMessage(img, messageBytes)
	if newImg == nil {
		return NewError(ErrEncoding, ErrMsgEncoding, nil)
	}

	// Encode and save image
	if err := png.Encode(outFile, newImg); err != nil {
		return NewError(ErrImageProcessing, ErrMsgImageProcessing, err)
	}

	return nil
}

// DecodeMessage extracts a message from an image using LSB steganography
func DecodeMessage(inputFile string) (string, error) {
	// Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		return "", NewError(ErrFileOperation, ErrMsgFileOperation, err)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return "", NewError(ErrImageProcessing, ErrMsgImageProcessing, err)
	}

	// Extract message
	extractedBytes := extractHiddenMessage(img)
	if len(extractedBytes) == 0 {
		return "", NewError(ErrDecoding, ErrMsgDecoding, nil)
	}

	return string(extractedBytes), nil
}

// hideMessage hides a message in an image using LSB
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
				r = uint32((uint8(r) & 0xFE) | bit)

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

// extractHiddenMessage extracts a hidden message from an image using LSB
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
