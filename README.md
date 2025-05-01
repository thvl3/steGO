# steGo

[![Go Report Card](https://goreportcard.com/badge/github.com/thule/steGo)](https://goreportcard.com/report/github.com/thule/steGo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/thule/steGo)](https://github.com/thule/steGo/releases)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/thule/steGo/build.yml?branch=main)](https://github.com/thule/steGo/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/thule/steGo.svg)](https://pkg.go.dev/github.com/thule/steGo)

A powerful command-line tool for hiding messages in images using steganography. steGo supports encoding and decoding messages in PNG images using the LSB (Least Significant Bit) method.

## Features

- **Message Encoding**: Hide messages in PNG images
- **Message Decoding**: Extract hidden messages from images
- **File Support**: Encode/Decode text files
- **Configuration**: Customizable through config files and environment variables
- **Shell Integration**: Shell completion for bash, zsh, fish, and PowerShell
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/thule/steGo.git
   cd steGo
   ```

2. Build the project:
   ```bash
   go build -o steGo ./cmd/steGo
   ```

3. Install the binary:
   ```bash
   # Linux/macOS
   sudo install steGo /usr/local/bin

   # Windows
   # Copy steGo.exe to a directory in your PATH
   ```

### From Release

Download the latest release from the [releases page](https://github.com/thule/steGo/releases) and install the appropriate binary for your platform.

## Usage

### Basic Usage

```bash
# Encode a message into an image
steGo encode -i input.png -o output.png -m "Secret message"

# Encode a message from a file
steGo encode -i input.png -o output.png -f message.txt

# Decode a message from an image
steGo decode -i output.png

# Decode a message to a file
steGo decode -i output.png -o message.txt
```

### Shell Completion

```bash
# Bash
source <(steGo completion bash)

# Zsh
steGo completion zsh > "${fpath[1]}/_steGo"

# Fish
steGo completion fish | source

# PowerShell
steGo completion powershell | Out-String | Invoke-Expression
```

### Configuration

steGo can be configured using a YAML configuration file. By default, it looks for `steGo.yaml` in:
- Linux/macOS: `$HOME/.config/steGo.yaml`
- Windows: `%APPDATA%\steGo\steGo.yaml`

Example configuration:
```yaml
verbose: true
default:
  input: "input.png"
  output: "output.png"
```

You can also use environment variables:
```bash
export STEGO_VERBOSE=true
export STEGO_DEFAULT_INPUT="input.png"
export STEGO_DEFAULT_OUTPUT="output.png"
```

## Command Reference

### Global Flags

- `--config string`: config file (default is $HOME/.config/steGo.yaml)
- `-v, --verbose`: verbose output
- `--version`: version for steGo

### Encode Command

```bash
steGo encode [flags]
```

Flags:
- `-f, --file string`: file containing message to encode
- `-h, --help`: help for encode
- `-i, --input string`: input image file (required)
- `-m, --message string`: message to encode
- `-o, --output string`: output image file (required)

### Decode Command

```bash
steGo decode [flags]
```

Flags:
- `-h, --help`: help for decode
- `-i, --input string`: input image file (required)
- `-o, --output string`: output file (optional)

## Development

### Prerequisites

- Go 1.21 or later
- Git

### Building

```bash
# Build for current platform
go build -o steGo ./cmd/steGo

# Cross-compile for all platforms
GOOS=linux GOARCH=amd64 go build -o steGo-linux-amd64 ./cmd/steGo
GOOS=darwin GOARCH=amd64 go build -o steGo-darwin-amd64 ./cmd/steGo
GOOS=windows GOARCH=amd64 go build -o steGo-windows-amd64.exe ./cmd/steGo
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
