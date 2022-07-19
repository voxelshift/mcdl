package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/voxelshift/mcdl/util"
)

var rootCmd = &cobra.Command{
	Use:     "mcdl",
	Short:   "mcdl is a simple tool to download various Minecraft server implementations",
	Version: util.Version,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		if debug, err := cmd.Flags().GetBool(debugFlag); debug && err == nil {
			log.SetLevel(log.DebugLevel)
		}
	},
}

var debugFlag = "debug"
var outputFlag = "output"

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP(debugFlag, "d", false, "show debug output")
	rootCmd.PersistentFlags().BoolP(outputFlag, "O", false, "set the output file")

	rootCmd.AddCommand(fabricCmd)
}
