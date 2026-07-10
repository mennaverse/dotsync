package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mennaverse/dotsync/apps/backend/manager"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [up|down]",
	Short: "Migrate the database schema to the latest version",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		direction := args[0]
		if direction != "up" && direction != "down" {
			fmt.Println("Invalid argument. Use 'up' or 'down'.")
			return
		}

		err := migrateDatabase(direction)
		if err != nil {
			fmt.Printf("Error migrating database: %v\n", err)
		} else {
			fmt.Printf("Database migrated %s successfully.\n", direction)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func getDatabase() (*sql.DB, error) {
	secrets := manager.NewSecretManager()
	dsn, err := secrets.GetDatabaseDSN()
	if err != nil {
		return nil, fmt.Errorf("failed to get database DSN: %w", err)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func migrateDatabase(direction string) error {
	db, err := getDatabase()
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	switch direction {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			return fmt.Errorf("failed to migrate database up: %w", err)
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			return fmt.Errorf("failed to migrate database down: %w", err)
		}
	}

	return nil
}
