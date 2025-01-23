package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/supakornn/reverse-shell-gen/payloads"
)

func ExecuteReverseShell(ip, port, shellType string) (string, error) {
	switch shellType {
	case "bash":
		return payloads.BashReverseShell(ip, port), nil
	case "python":
		return payloads.PythonReverseShell(ip, port), nil
	case "php":
		return payloads.PhpReverseShell(ip, port), nil
	case "powershell":
		return payloads.PowerShellReverseShell(ip, port), nil
	case "netcat":
		return payloads.NetcatReverseShell(ip, port), nil
	default:
		return "", errors.New("unsupported shell type")
	}
}

func ListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all supported reverse shell types",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Supported types:")
			fmt.Println("- bash")
			fmt.Println("- python")
			fmt.Println("- php")
			fmt.Println("- powershell")
			fmt.Println("- netcat")
		},
	}
}
