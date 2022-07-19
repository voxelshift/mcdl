# mcdl

mcdl is a simple cli tool to download various Minecraft server implementations from their official sources.

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
