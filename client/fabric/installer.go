package fabric

import (
	"github.com/voxelshift/mcdl/client"
	"github.com/voxelshift/mcdl/util"
)

var GetInstallerVersionsUrl = "https://meta.fabricmc.net/v2/versions/installer"

type InstallerVersions []InstallerVersion

type InstallerVersion struct {
	Name   string `json:"version"`
	Build  int
	Stable bool
}

func GetInstallerVersions() (*InstallerVersions, error) {
	resp := &InstallerVersions{}
	err := client.GetJson(GetInstallerVersionsUrl, resp)

	return resp, err
}

// Get the latest (possibly unstable) installer version
func (versions InstallerVersions) GetLatest() *InstallerVersion {
	return &versions[0]
}

// Get the latest stable installer version
func (versions InstallerVersions) GetStable() *InstallerVersion {
	for _, v := range versions {
		if v.Stable {
			return &v
		}
	}

	return nil
}

// Get the installer version that matches the passed string
func (versions InstallerVersions) GetVersion(name string) *InstallerVersion {
	if name == util.LatestVersion {
		return versions.GetLatest()
	}

	if name == util.StableVersion {
		return versions.GetStable()
	}

	for _, v := range versions {
		if name == v.Name {
			return &v
		}
	}

	return nil
}
