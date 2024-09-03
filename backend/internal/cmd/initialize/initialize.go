package initialize

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Init the data, config...etc",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
}

func init() {
}
