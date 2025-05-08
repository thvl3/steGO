# steGo

[![Go Report Card](https://goreportcard.com/badge/github.com/thvl3/steGo)](https://goreportcard.com/report/github.com/thvl3/steGo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/thvl3/steGo)](https://github.com/thvl3/steGo/releases)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/thvl3/steGo/build.yml?branch=main)](https://github.com/thvl3/steGo/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/thvl3/steGo.svg)](https://pkg.go.dev/github.com/thvl3/steGo)

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
   git clone https://github.com/thvl3/steGo.git
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

   # Windows (PowerShell)
   # Create a directory for your executables if you don't have one, e.g., $HOME\bin
   # mkdir $HOME\bin -ErrorAction SilentlyContinue
   # Copy steGo.exe to that directory:
   Copy-Item .\steGo.exe -Destination "$HOME\bin\steGo.exe"
   # Ensure that the directory (e.g., $HOME\bin) is in your PATH.
   # You can check by running: $env:Path -split ';'
   # To add it temporarily for the current session: $env:Path += ";$HOME\bin"
   # For a permanent change, search for "Edit the system environment variables" in Windows.
   ```

### Windows Execution Notes

Sometimes, antivirus software on Windows might incorrectly flag `steGo.exe` as malicious. If you encounter this, or face permission issues when trying to run the program, consider the following:

*   **Antivirus Exception**: Add an exception for `steGo.exe` in your antivirus software. The steps for this vary depending on your antivirus program.
*   **Run as Administrator (Use with Caution)**: You can try running PowerShell as an administrator. Right-click on the PowerShell icon and select "Run as administrator". However, be cautious when running any program with elevated privileges.
*   **Verify the Source**: Ensure you have downloaded `steGo` from the official GitHub releases page to minimize risks.

### From Release

Download the latest release from the [releases page](https://github.com/thvl3/steGo/releases) and install the appropriate binary for your platform:

| Platform | Binary Name |
|----------|------------|
| Linux    | `steGo-linux-amd64` |
| macOS    | `steGo-darwin-amd64` |
| Windows  | `steGo-windows-amd64.exe` |

After downloading, rename the binary to `steGo` (or `steGo.exe` on Windows) and install it in your PATH.

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
