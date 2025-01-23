# 🔄 Reverse Shell Generator

A simple tool to generate reverse shell commands for various programming languages and shells.

## ✨ Features

- 🛠️ Multiple language support
- 🔌 Easy-to-use command generation
- 🔒 Base64 encoding option
- 📝 Copy to clipboard functionality

## 🔧 Installation

```bash
git clone https://github.com/yourusername/reverse-shell-gen
cd reverse-shell-gen
go build
./reverse-shell-gen
```

## 🚀 Usage

List all available shell types:

```bash
./reverse-shell-gen list
```

Generate a reverse shell command:

```bash
./reverse-shell-gen --ip <IP_ADDRESS> --port <PORT> --type <SHELL_TYPE>
```

Example:

```bash
./reverse-shell-gen --ip 192.168.1.10 --port 4444 --type bash
```

Available shell types can be viewed using the `list` command.
