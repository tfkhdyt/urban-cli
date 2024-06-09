[![ReadMeSupportPalestine](https://raw.githubusercontent.com/Safouene1/support-palestine-banner/master/banner-project.svg)](https://github.com/Safouene1/support-palestine-banner)

# Urban CLI

Urban CLI is blazingly fast command line interface for Urban Dictionary written
in Go. This tool allows users to quickly search for definitions and slang terms
from the Urban Dictionary directly from their terminal. Its speed and simplicity
make it an ideal choice for developers and tech enthusiasts who prefer command
line tools.

## Key Features

- **Blazing Fast Performance**: Leveraging the efficiency of the Go programming
  language, Urban CLI delivers rapid search results.
- **Command Line Interface**: Designed for users who prefer working in the
  terminal, offering a straightforward and efficient way to access Urban
  Dictionary.
- **Easy Installation**: Simple installation process with precompiled binaries
  available for various operating systems.
- **Cross-Platform Support**: Compatible with major operating systems including
  Linux, macOS, and Windows.
- **Open Source**: Freely available source code under an open-source license,
  encouraging community contributions and transparency.

## Getting Started

### Installation

- **Build from source**: `go install github.com/tfkhdyt/urban-cli`
- **Arch Linux (AUR)**: `yay -S urban-cli-bin`
- **Standalone binary:** Download the binary file from
  [release page](https://github.com/tfkhdyt/urban-cli/releases) and move the
  binary to one of the `PATH` directories in your system, for example:
  - **Linux:** `$HOME/.local/bin/`, `/usr/local/bin/`
  - **Windows:** `%LocalAppData%\Programs\`
  - **macOS:** `/usr/local/bin/`

### Usage

```bash
# find definition
urban-cli rizz

# find multi words definition using quote
urban-cli 'kai cenat'

# limit entries (default: 7)
urban-cli skibidi --max 15 # or -m 15

# reverse entries (from bottom to top)
urban-cli mewing --reverse # or -r
```

More details in `urban-cli --help`

## License

This project is licensed under the GPLv3 License. See the LICENSE file for more
details.
