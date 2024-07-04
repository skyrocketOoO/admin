package run

import (
	"admin/internal"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Long:  ``,
	Run:   internal.RunServer,
}

func init() {
	RunCmd.Flags().StringP("port", "p", ":50051", "port")
	RunCmd.Flags().
		StringP("database", "d", "sqlite", `database enum. allowed: "postgres", "sqlite"`)
}
