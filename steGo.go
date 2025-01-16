package main


import (
  "bytes"
  "fmt"
  "image"
  "image/color"
  "image/png"
  "os"
  "strings"
)

const delimiter = "##END##"


// Convert input string to binary for encoding
// Private helper function for encodeMessage function
func stringToBinary(s string) string {
  var binaryString string
  for _, c := range s {
    binaryString += fmt.Sprintf("%08b", c)
  }
  return binaryString
}

func binaryToString(binary string) string {
  var text string
  for i := 0; i < len(binary); i += 8 {
    if i+8 > len(binary){
      break
    }
    var char byte
    fmt.Sscanf(binary[i:i+8], "%b", &char)
    text += string(char)
  }
  return text
}
// Encode binary message into file
func encodeMessage(inputFile, outputFile, message string) error {
  //open input image

  file, err := os.Open(inputFile)
  if err != nil {
    return err
  }
  defer file.Close()

  img, err := png.Decode(file)
  if err != nil {
    return err
  }

  //Convert message to binary and add a delimiter

  binaryMessage := stringToBinary(message + delimiter)

  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  //Create the new image with the encoded and modified pixels

  newImg := image.NewRGBA(bounds)

  bitIndex := 0
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      r, g, b, a := img.At(x, y).RGBA()

      if bitIndex < len(binaryMessage) {
        //Modify the LSB of the red channel
        r = (r &^ 1) | uint32(binaryMessage[bitIndex]-'0')
        bitIndex++
      }

      newImg.Set(x,y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
    }
  }

  //Save the modified image
  outFile, err := os.Create(outputFile)
  if err != nil {
    return err
  }
  defer outFile.Close()

  return png.Encode(outFile, newImg)
}

//Decode a message from a stego image

func decodeMessage(inputFile string)(string, error){
  //open inout image
  file, err := os.Open(inputFile)
  if err != nil {
    return "", err
  
  }
  defer file.Close()

  img, err := png.Decode(file)
  if err != nil {
    return "", err
  }

  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  var binaryMessage bytes.Buffer

  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      r, _, _, _ := img.At(x, y).RGBA()
      binaryMessage.WriteString(fmt.Sprintf("%b", r&1))

      //Check if we have reached Delim string
      if strings.Contains(binaryToString(binaryMessage.String()), delimiter) {
        message := binaryToString(binaryMessage.String())
        return strings.Split(message, delimiter)[0], nil
      }
    }
  }
  return "", fmt.Errorf("No hidden message found")
}

// CLI Interface

func main() {
  if len(os.Args) <  5 {
    fmt.Println("Usage:")
    fmt.Println(" Encoding: ./stego -mode encode -input <inputfile.png> -output output.png -message \"secret message\"")
    fmt.Println(" Decoding: ./stego -mode decode -input <encodedfile.png>")
    return
  }

  mode := os.Args[2]
  inputFile := os.Args[4]

  switch mode {

  case "encode":
    if len(os.Args) < 8 {
      fmt.Println("Missing arguments for encoding.")
      return
    }
    outputFile := os.Args[6]
    message := os.Args[8]
    err := encodeMessage(inputFile, outputFile, message)
    if err != nil {
      fmt.Println("Encoding error: ", err)
      return
    }
    fmt.Println("Message successfully encoded!")
  case "decode":
    message, err := decodeMessage(inputFile)
    if err != nil {
      fmt.Println("Decoding error: ", err)
      return
    }
    fmt.Println("Decoded Message:", message)

  default:
    fmt.Println("Invalid mdoe, please use 'encode' or 'decode'.")
  }
}
