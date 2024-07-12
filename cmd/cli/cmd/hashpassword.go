/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// hashpasswordCmd represents the hashpassword command
var hashpasswordCmd = &cobra.Command{
	Use:   "hash",
	Short: "Displays the hash of a password",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	RunE:  HashPassword,
}

func HashPassword(cmd *cobra.Command, args []string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(args[0]), 14)
	if err != nil {
		return err
	}

	fmt.Printf("Hash of %q = %q\n", args[0], passwordHash)
	return nil
}

func init() {
	rootCmd.AddCommand(hashpasswordCmd)
}
