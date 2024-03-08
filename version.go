package main

import (
	"fmt"
	"runtime/debug"
)

var Version = "dev"

func init() {
	info, ok := debug.ReadBuildInfo()

	if ok {
		fmt.Println("Build Info", info.Main.Path, info.Main.Version)
	}
}
