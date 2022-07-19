package fabric

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/voxelshift/mcdl/client"
)

var DownloadFabricUrl = "https://meta.fabricmc.net//v2/versions/loader/%s/%s/%s/server/jar"

type DownloadFabricOptions struct {
	GameVersion      string
	LoaderVersion    string
	InstallerVersion string
	Output           string
}

func DownloadFabric(options DownloadFabricOptions) (*string, error) {
	gameVersions, err := GetGameVersions()
	if err != nil {
		return nil, err
	}
	gameVersion := gameVersions.GetVersion(options.GameVersion).Name

	loaderVersions, err := GetLoaderVersions(gameVersion)
	if err != nil {
		return nil, err
	}
	loaderVersion := loaderVersions.GetVersion(options.LoaderVersion).Name

	installerVersions, err := GetInstallerVersions()
	if err != nil {
		return nil, err
	}
	installerVersion := installerVersions.GetVersion(options.InstallerVersion).Name

	output := &options.Output
	if *output == "" {
		output = nil
	}

	log.Infof("downloading Fabric %s (loader: %s; installer: %s)", gameVersion, loaderVersion, installerVersion)

	url := fmt.Sprintf(DownloadFabricUrl, gameVersion, loaderVersion, installerVersion)
	return client.Download(url, output)
}
