/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var userName string

// listcampaignsCmd represents the listcampaigns command
var listcampaignsCmd = &cobra.Command{
	Use:   "listcampaigns",
	Short: "Lists the campaigns",
	Long:  ``,
	RunE:  ListCampaigns,
}

func ListCampaigns(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	queries, err := initializeDb()
	if err != nil {
		return err
	}

	id, err := queries.GetUserId(ctx, userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("User with the username of %q not found\n", userName)
			return nil
		} else {
			return err
		}
	}

	campaigns, err := queries.GetCampaigns(ctx, id.(int64))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("User %q has no campaigns\n", userName)
			return nil
		} else {
			return err
		}
	}

	for i, c := range campaigns {
		fmt.Printf("%d: %s [%s]\n", i+1, c.Name, c.KeyField)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listcampaignsCmd)

	listcampaignsCmd.PersistentFlags().StringVarP(&userName, "user", "u", "harry@culpan.org", "Username to view campaigns for")
}
