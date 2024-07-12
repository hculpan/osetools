/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	_ "github.com/glebarez/go-sqlite"

	"github.com/hculpan/osetools/internal/dbutils"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Create an empty DB",
	Long:  ``,
	RunE:  InitDB,
}

func InitDB(cmd *cobra.Command, args []string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing .env file: %w", err)
	}

	// Update this whatever variables you need to retrieve to
	// initialize the db
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return errors.New("missing DB_PATH in configuration file")
	}

	if _, err := os.Stat(dbPath); err == nil {
		err := os.Remove(dbPath)
		if err != nil {
			return err
		}
	}
	os.MkdirAll(filepath.Dir(dbPath), 0755)
	os.Create(dbPath)

	dbase, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer dbase.Close()

	if _, err := dbase.ExecContext(context.Background(), dbutils.GetDdl()); err != nil {
		return err
	}

	fmt.Println("New DB created")

	return err
}

func init() {
	rootCmd.AddCommand(initdbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
