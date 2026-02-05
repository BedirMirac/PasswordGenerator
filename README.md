# pwgen
<div align="center">

**A simple yet robust command-line tool for generating strong, random passwords.**

</div>

`pwgen` is a simple and lightweight **command-line password generator written in Go (Golang)**.  
It allows you to quickly generate secure and random passwords directly from the terminal.

This project focuses on simplicity, readability, and ease of use.

---

## Features

- Secure random password generation
- Fast and minimal CLI tool
- Written in pure Go (no external dependencies)
- Easy to understand and extend
- Cross-platform support

---

## Requirements

- Go 1.18 or later

Check your Go version:

```bash
go version
```
## Installation
### Option 1: Install using go install (recommended)

#### You can install pwgen globally using Go:

```
go install github.com/BedirMirac/PasswordGenerator@latest
```

#### Make sure your Go bin directory is in your PATH:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

#### After installation, you can run the tool from anywhere:

```
pwgen
```

### Option 2: Build from source

#### Clone the repository:
```
git clone https://github.com/BedirMirac/PasswordGenerator.git
cd PasswordGenerator
```
#### Build the project:
```
go build
```
#### Run the executable:
```
./pwgen
```
## Usage

After building or installing, you can run the `pwgen` command.

### Basic Command

Generate a password with default settings (e.g., a mix of character types and a default length):
```bash
# If built locally
./pwgen

# If installed globally
pwgen
```

### Available Options

Use flags to customize your password:

| Flag          | Alias | Description                                    | Default   |
|---------------|-------|------------------------------------------------|-----------|
| `--length`    | `-l`  | Specify the length of the password.            | `16`      |
| `--numbers`   | `-d`  | Include numbers (0-9).               | `5`    |
| `--symbols`   | `-c`  | Include special symbols (`!@#$%^&*()-+=[]{}<>`).               | `5`    |
| `--vault`   | `-v`  | To see the your password vault.                         | `false`    |
| `--id`   | `-i`  | To udpate or remove a password you need to enter an id so that pwgen can work successfully. | `0`   |
| `--udpate`   | `-u`  | To update a password that exists. | `false`    |
| `--remove`   | `-r` | To remove a password that exists. | `false`    |
| `--auto`|   `-a` | To change  password with a random generated password | `false` |
| `--new-pass` | `-p` | To change password with a password that you created | `" "`|
| `--save` | `-s` | To save the generated password | `false`|
| `--name` | `-n` | To name the password generated | Date and Time when password created |

_Note: If no character types are explicitly enabled, the tool will default to including all types for a robust password._

### Examples

```bash
# Generate a 20-character password with all character types
./pwgen --length 20

# Generate a very short, number-only PIN
./pwgen -l 4 -d 4

# Remove the password whose id 5
./pwgen -r -i 5
```

## 📁 Project Structure

```
PasswordGenerator/
├── .gitignore         # Specifies intentionally untracked files to ignore
├── LICENSE            # MIT License for the project
├── README.md          # Project README file
├── cmd/               # Main logic, command definitions, and DB operations
├── go.mod             # Go module definition file
├── go.sum             # Go module checksums for dependencies
└── main.go            # Main entry point for the password generator application
```


## 🤝 Contributing

We welcome contributions! If you have suggestions for improvements, new features, or bug fixes, please open an issue or submit a pull request.

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

## 🙏 Acknowledgments

-   Built with the power of the Go programming language.

## 📞 Support & Contact

-   🐛 Issues: [GitHub Issues](https://github.com/BedirMirac/PasswordGenerator/issues)

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

Made with ❤️ by [BedirMirac](https://github.com/BedirMirac)

</div>