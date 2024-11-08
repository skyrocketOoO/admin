package run

import (
	"admin/internal"
	"admin/internal/service/orm"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Long:  ``,
	Run:   internal.RunServer,
}

func init() {
	Cmd.Flags().BoolVar(&orm.Debug, "debug", false, "Enable debug mode")
	Cmd.Flags().StringP("port", "p", ":50051", "port")
	Cmd.Flags().
		StringP("database", "d", "sqlite", `database enum. allowed: "postgres", "sqlite"`)
}
