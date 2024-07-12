/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/hculpan/osetools/internal/db"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	RunE:  AddUser,
}

func AddUser(cmd *cobra.Command, args []string) error {
	username := args[0]
	password := args[1]
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	queries, err := initializeDb()
	if err != nil {
		return err
	}

	user, err := queries.InsertUser(context.Background(), db.InsertUserParams{
		Username: username,
		Password: string(passwordHash),
	})
	if err != nil {
		return err
	}

	fmt.Printf("User %s inserted\n", user.Username)

	closeDb()

	return nil
}

func init() {
	rootCmd.AddCommand(adduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
