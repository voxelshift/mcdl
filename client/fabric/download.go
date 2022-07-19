package fabric

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/voxelshift/mcdl/client"
)

var DownloadFabricUrl = "https://meta.fabricmc.net//v2/versions/loader/%s/%s/%s/server/jar"

var (
	errInvalidGameVersion      = errors.New("game version not found")
	errInvalidLoaderVersion    = errors.New("loader version not found")
	errInvalidInstallerVersion = errors.New("installer version not found")
)

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
	gameVersion := gameVersions.GetVersion(options.GameVersion)
	if gameVersion == nil {
		return nil, errInvalidGameVersion
	}

	loaderVersions, err := GetLoaderVersions(gameVersion.Name)
	if err != nil {
		return nil, err
	}
	loaderVersion := loaderVersions.GetVersion(options.LoaderVersion)
	if loaderVersion == nil {
		return nil, errInvalidLoaderVersion
	}

	installerVersions, err := GetInstallerVersions()
	if err != nil {
		return nil, err
	}
	installerVersion := installerVersions.GetVersion(options.InstallerVersion)
	if installerVersion == nil {
		return nil, errInvalidInstallerVersion
	}

	output := &options.Output
	if *output == "" {
		output = nil
	}

	log.Infof("downloading Fabric %s (loader: %s; installer: %s)", gameVersion.Name, loaderVersion.Name, installerVersion.Name)

	url := fmt.Sprintf(DownloadFabricUrl, gameVersion.Name, loaderVersion.Name, installerVersion.Name)
	return client.Download(url, output)
}
