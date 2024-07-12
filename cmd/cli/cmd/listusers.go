/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// listusersCmd represents the listusers command
var listusersCmd = &cobra.Command{
	Use:   "listusers",
	Short: "Lists registered users",
	Long:  ``,
	RunE:  ListUsers,
}

func ListUsers(cmd *cobra.Command, args []string) error {
	queries, err := initializeDb()
	if err != nil {
		return err
	}

	Users, err := queries.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(Users) == 0 {
		fmt.Println("There are no users")
	} else {
		for _, user := range Users {
			fmt.Printf("User %d: %s, %s, %s", user.ID, user.Username, user.Realname.String, user.CreateDatetime.String())
		}
	}

	closeDb()

	return nil
}

func init() {
	rootCmd.AddCommand(listusersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listusersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listusersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
