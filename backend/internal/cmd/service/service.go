package service

import (
	"admin/internal/cmd/service/run"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "service",
	Short: "The main service command",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(run.Cmd)
}
