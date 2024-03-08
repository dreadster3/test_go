package main

import (
	"runtime/debug"
)

var Version = "dev"

func init() {
	if Version == "dev" {
		info, ok := debug.ReadBuildInfo()

		if ok {
			Version = info.Main.Version
		}
	}
}
