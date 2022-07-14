package fabric

import "github.com/voxelshift/mcdl/util"

var GetGameVersionsUrl = "https://meta.fabricmc.net/v2/versions/game"

type GameVersion struct {
	Name   string
	Stable bool
}

type GameVersions []GameVersion

func GetGameVersions() (GameVersions, error) {
	resp := GameVersions{}
	err := util.GetJson(GetGameVersionsUrl, resp)

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
	for _, v := range versions {
		if name == v.Name {
			return &v
		}
	}

	return nil
}
