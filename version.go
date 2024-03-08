package main

import (
	"runtime/debug"
)

var Version = "dev"

func init() {
	// Used when using go install
	if Version == "dev" {
		info, ok := debug.ReadBuildInfo()

		if ok {
			Version = info.Main.Version
		}
	}
}
