package cmd

import (
	"fmt"

	"github.com/notblessy/config"

	"github.com/spf13/cobra"
)

var runMigration = &cobra.Command{
	Use:   "migration",
	Short: "run migration",
	Long:  `This subcommand is for execute database migration`,
	Run:   migrate,
}

func init() {
	rootCmd.AddCommand(runMigration)
	config.LoadConfig()
}

func migrate(cmd *cobra.Command, args []string) {
	fmt.Println("Migration")
}
