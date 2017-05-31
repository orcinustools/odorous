package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/docker/go-plugins-helpers/volume"
)

const socketAddress = "/run/docker/plugins/odorous.sock"
const glusterfsID = "_glusterfs"

var (
	defaultDir  = filepath.Join(volume.DefaultDockerRootDirectory, glusterfsID)
	root        = flag.String("root", defaultDir, "GlusterFS volumes root directory")
)

func main() {
	d := newGlusterfsDriver(*root)
	h := volume.NewHandler(d)
	fmt.Println(h.ServeUnix(socketAddress, 0))
}
