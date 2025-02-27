/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"

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

	dbServer := os.Getenv("DB_SERVER")
	if dbServer == "" {
		log.Fatal("missing DB_SERVER in configuration file")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("missing DB_NAME in configuration file")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("missing DB_USER in configuration file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("missing DB_PASSWORD in configuration file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbServer, dbName)

	dbase, err = sql.Open("mysql", dsn)
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
