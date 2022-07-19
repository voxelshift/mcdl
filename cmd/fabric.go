package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/voxelshift/mcdl/client/fabric"
	"github.com/voxelshift/mcdl/util"
)

var fabricCmd = &cobra.Command{
	Use:   "fabric <version>",
	Short: "download server files from the Fabric project",
	Args:  cobra.MaximumNArgs(1),
	Run:   run,
}

var loaderFlag = "loader"
var installerFlag = "installer"

func run(cmd *cobra.Command, args []string) {
	var gameVersion string
	if len(args) == 1 {
		gameVersion = args[0]
	} else {
		gameVersion = util.StableVersion
	}

	var loaderVersion = util.UnwrapFlag(cmd.Flags().GetString(loaderFlag))
	var installerVersion = util.UnwrapFlag(cmd.Flags().GetString(installerFlag))

	var output, err = cmd.Flags().GetString(outputFlag)
	if err != nil {
		log.Debugf("failed to parse output flag: %v", err)
		output = ""
	}

	filename, err := fabric.DownloadFabric(fabric.DownloadFabricOptions{
		GameVersion:      gameVersion,
		LoaderVersion:    loaderVersion,
		InstallerVersion: installerVersion,
		Output:           output,
	})
	if err != nil {
		log.Panicf("failed to download fabric server")
		os.Exit(1)
	}
	log.Infof("successfully downloaded %s", *filename)
}

func init() {
	fabricCmd.Flags().StringP(loaderFlag, "l", util.StableVersion, "specify the Fabric loader version")
	fabricCmd.Flags().StringP(installerFlag, "i", util.StableVersion, "specify the Fabric installer version")
}
