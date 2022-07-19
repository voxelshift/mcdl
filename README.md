# mcdl

mcdl is a simple cli tool to download various Minecraft server implementations from their official sources.

## installation

### docker

mcdl has a Docker image hosted on GitHub Container Registry:

```Dockerfile
# You probably shouldn't use "latest" in production images since things may break after a major version bump.
# Instead you can pin to a specific version.
FROM ghcr.io/voxelshift/mcdl:latest

RUN mcdl fabric 1.19
```

### binaries

Linux binaries for mcdl are available on the [Releases](https://github.com/voxelshift/mcdl/releases/tag/0.0.1) page on
GitHub. The goal is to eventually get mcdl on package repositories when it becomes more stable and featureful.

Windows and Mac binaries aren't currently provided as I don't see there being much use but if there's enough demand for
that they can always be added.

## projects

### fabric

mcdl can be used to download [Fabric](https://fabricmc.net/) directly from the [Fabric Meta API](https://github.com/FabricMC/fabric-meta):

```sh
# Download the latest stable version of Fabric
mcdl fabric

# Download Fabric and specify output file
mcdl fabric -O=fabric.jar

# Download Fabric 1.19
mcdl fabric 1.19

# Download Fabric with a specific loader and installer version
mcdl fabric 1.19 -l=0.14.8 -i=0.10.2
```
