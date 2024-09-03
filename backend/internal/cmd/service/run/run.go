package run

import (
	"admin/internal"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Long:  ``,
	Run:   internal.RunServer,
}

func init() {
	Cmd.Flags().StringP("port", "p", ":50051", "port")
	Cmd.Flags().
		StringP("database", "d", "sqlite", `database enum. allowed: "postgres", "sqlite"`)
}
