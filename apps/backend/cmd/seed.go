package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/mennaverse/dotsync/apps/backend/manager"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := getDatabase()
		if err != nil {
			fmt.Printf("Error getting database: %v\n", err)
			return
		}
		defer db.Close()

		err = seedDatabase(db)
		if err != nil {
			fmt.Printf("Error seeding database: %v\n", err)
		} else {
			fmt.Println("Database seeded successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

type SeedUser struct {
	Username      string
	Email         string
	EmailVerified bool
	Password      string
}

func seedDatabase(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := seedUsers(tx); err != nil {
		return fmt.Errorf("failed to seed users: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func seedUsers(tx *sql.Tx) error {
	f, err := os.Open("seeders/users.csv")
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer f.Close()

	users := new([]SeedUser)
	if err := gocsv.UnmarshalFile(f, users); err != nil {
		return fmt.Errorf("failed to unmarshal CSV data: %w", err)
	}

	secrets := manager.NewSecretManager()
	cryptoManager := manager.NewCryptoManager(secrets)

	for _, user := range *users {
		passwordHash, err := cryptoManager.HashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		_, err = tx.Exec(
			"INSERT INTO users (username, email, email_verified, password) VALUES (?, ?, ?, ?)",
			user.Username, user.Email, user.EmailVerified, passwordHash,
		)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}

	return nil
}
