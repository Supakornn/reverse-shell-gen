# 🔄 Reverse Shell Generator

A simple tool to generate reverse shell commands for various programming languages and shells.

## ✨ Features

- 🛠️ Multiple language support
- 🔌 Easy-to-use command generation

## 🔧 Installation

```bash
git clone https://github.com/yourusername/reverse-shell-gen
cd reverse-shell-gen
go mod tidy
go build
./reverse-shell-gen
```

## 🚀 Usage

List all available shell types:

```bash
./reverse-shell-gen list
```

```bash
Supported shell types:
- bash
- python
- php
- powershell
- netcat
```

Generate a reverse shell command:

```bash
./reverse-shell-gen --ip <IP_ADDRESS> --port <PORT> --type <SHELL_TYPE>
```

Example:

```bash
./reverse-shell-gen --ip 192.168.1.10 --port 4444 --type bash
```
