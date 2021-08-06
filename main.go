package main

import (
	"flag"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var mountpoint = flag.String("m", "tillybon", "name of the mount")

func main() {
	flag.Parse()

	c, err := fuse.Mount(
		*mountpoint,
		fuse.FSName("tillybon"),
		fuse.Subtype("tillybon-fs"),
		fuse.LocalVolume(),
		fuse.VolumeName("your Go package"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	err = fs.Serve(c, FS{})
	if err != nil {
		log.Fatal(err)
	}

	// check if the mount process has an error to report
	<-c.Ready
	if err := c.MountError; err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(*mountpoint); err != nil {
		log.Fatal(err)
	}
}
