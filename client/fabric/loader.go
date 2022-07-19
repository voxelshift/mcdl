package fabric

import (
	"fmt"

	"github.com/voxelshift/mcdl/client"
	"github.com/voxelshift/mcdl/util"
)

var GetLoaderVersionsUrl = "https://meta.fabricmc.net/v2/versions/loader/%s"

type LoaderVersions []LoaderVersion

type LoaderVersion struct {
	Name   string `json:"version"`
	Build  int
	Stable bool
}

type GetLoaderVersionsResponse struct {
	Loader LoaderVersion

	// ...there are other fields but they are irrelevant for this...
}

func GetLoaderVersions(gameVersion string) (*LoaderVersions, error) {
	url := fmt.Sprintf(GetLoaderVersionsUrl, gameVersion)

	resp := &[]GetLoaderVersionsResponse{}
	err := client.GetJson(url, resp)
	if err != nil {
		return nil, err
	}

	versions := make(LoaderVersions, len(*resp))
	for i, version := range *resp {
		versions[i] = version.Loader
	}

	return &versions, nil
}

// Get the latest (possibly unstable) loader version
func (versions LoaderVersions) GetLatest() *LoaderVersion {
	return &versions[0]
}

// Get the latest stable loader version
func (versions LoaderVersions) GetStable() *LoaderVersion {
	for _, v := range versions {
		if v.Stable {
			return &v
		}
	}

	return nil
}

// Get the loader version that matches the passed string
func (versions LoaderVersions) GetVersion(name string) *LoaderVersion {
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
