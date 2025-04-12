/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/hculpan/osetools/internal/db"
	"github.com/spf13/cobra"
)

// syncxpCmd represents the syncxp command
var syncxpCmd = &cobra.Command{
	Use:   "syncxp",
	Short: "Update all characters's total_xp",
	Long: `For each character, this will sum the xp awards and update
	the character table with the total. This should only be necessary
	the first time after the update is deployed, but it might be useful 
	later too, who knows?`,
	RunE: SyncXp,
}

func SyncXp(cmd *cobra.Command, args []string) error {
	queries, err := initializeDb()
	if err != nil {
		return err
	}
	defer closeDb()

	fmt.Printf("Updating all characters with XP Award totals\n")

	characters, err := queries.GetAllCharacters(context.Background())
	if err != nil {
		return err
	}

	for _, c := range characters {
		xpAwards, err := queries.GetXpAwardsForCharacter(context.Background(), c.ID)
		if err != nil {
			return err
		}

		totalXp := 0
		for _, xp := range xpAwards {
			totalXp += int(xp.XpAwardWithBonus)
		}
		fmt.Printf("  Updating %s with %d\n", c.Name, totalXp)
		queries.UpdateCharacter(context.Background(), db.UpdateCharacterParams{
			Name:       c.Name,
			PlayerName: c.PlayerName,
			XpBonus:    c.XpBonus,
			TotalXp:    int32(totalXp),
			ID:         c.ID,
		})
	}

	fmt.Println("Done")

	return nil
}

func init() {
	rootCmd.AddCommand(syncxpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncxpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncxpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
