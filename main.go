package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/supakornn/reverse-shell-gen/cmd"
)

func main() {
	var ip, port, shellType string

	rootCmd := &cobra.Command{
		Use:   "reverse-shell-gen",
		Short: "Reverse Shell Generator CLI",
		Run: func(c *cobra.Command, args []string) {
			payload, err := cmd.ExecuteReverseShell(ip, port, shellType)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Generated Payload:")
			fmt.Println(payload)
		},
	}

	rootCmd.Flags().StringVar(&ip, "ip", "127.0.0.1", "IP address for reverse shell")
	rootCmd.Flags().StringVar(&port, "port", "4444", "Port for reverse shell")
	rootCmd.Flags().StringVar(&shellType, "type", "bash", "Type of reverse shell (e.g., bash, python)")

	rootCmd.AddCommand(cmd.ListCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
