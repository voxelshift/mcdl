package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/voxelshift/mcdl/util"
)

var rootCmd = &cobra.Command{
	Use:     "mcdl [flags] [command]",
	Short:   "mcdl is a simple tool to download various Minecraft server implementations",
	Run:     func(cmd *cobra.Command, args []string) {},
	Version: util.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}
