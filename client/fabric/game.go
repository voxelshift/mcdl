package fabric

import (
	"github.com/voxelshift/mcdl/client"
	"github.com/voxelshift/mcdl/util"
)

var GetGameVersionsUrl = "https://meta.fabricmc.net/v2/versions/game"

type GameVersion struct {
	Name   string `json:"version"`
	Stable bool
}

type GameVersions []GameVersion

func GetGameVersions() (*GameVersions, error) {
	resp := &GameVersions{}
	err := client.GetJson(GetGameVersionsUrl, resp)

	return resp, err
}

// Get the latest (possibly unstable) game version
func (versions GameVersions) GetLatest() *GameVersion {
	return &versions[0]
}

// Get the latest stable game version
func (versions GameVersions) GetStable() *GameVersion {
	for _, v := range versions {
		if v.Stable {
			return &v
		}
	}

	return nil
}

// Get the game version that matches the passed string
func (versions GameVersions) GetVersion(name string) *GameVersion {
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
