# ODOROUS - Docker volume plugin for GlusterFS

[![Build Status](https://travis-ci.org/orcinustools/odorous.svg?branch=master)](https://travis-ci.org/orcinustools/odorous)

## Usage

1 - Install the plugin

```
$ docker plugin install orcinus/odorous:0.1.1 # or docker plugin install vieux/sshfs DEBUG=1
```

2 - Create a volume

```
$ docker volume create -d orcinus/odorous:0.1.1 -o servers=<List of glusterfs servers> [-o res=<URL to glusterfsrest api>] [-o source=<Base directory where volumes are created in the cluster>] glustervolume

$ docker volume create -d orcinus/odorous:next -o servers=gfs-1:gfs-2:gfs-3 -o rest=http://gfs-1:9000 -o source=/var/lib/gluster/volumes glustervolumes
$ docker volume ls
DRIVER              VOLUME NAME
local               2d75de358a70ba469ac968ee852efd4234b9118b7722ee26a1c5a90dcaea6751
local               842a765a9bb11e234642c933b3dfc702dee32b73e0cf7305239436a145b89017
local               9d72c664cbd20512d4e3d5bb9b39ed11e4a632c386447461d48ed84731e44034
local               be9632386a2d396d438c9707e261f86fd9f5e72a7319417901d84041c8f14a4d
local               e1496dfe4fa27b39121e4383d1b16a0a7510f0de89f05b336aab3c0deb4dda0e
orcinus/odorous 	glustervolume
```

3 - Use the volume

```
$ docker run -it -v glustervolume:<path> busybox ls <path>
```

## LICENSE

MIT
