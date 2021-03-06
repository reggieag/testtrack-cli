package cmds

import (
	"github.com/Betterment/testtrack-cli/migrationrunners"
	"github.com/Betterment/testtrack-cli/servers"
	"github.com/spf13/cobra"
)

var migrateDoc = `
Runs all migrations that haven't been applied to the TestTrack server yet.
`

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run outstanding migrations",
	Long:  migrateDoc,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return migrate()
	},
}

func migrate() error {
	server, err := servers.New()
	if err != nil {
		return err
	}
	runner, err := migrationrunners.New(server)
	if err != nil {
		return err
	}

	err = runner.RunOutstanding()
	if err != nil {
		return err
	}

	return nil
}
