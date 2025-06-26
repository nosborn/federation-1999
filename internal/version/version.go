package version

import (
	_ "embed"
	"runtime/debug"
	"strings"
)

//go:embed VERSION
var baseVersion string

func String() string {
	version := strings.TrimSpace(baseVersion)
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" {
				return version + "+" + s.Value[:7]
			}
		}
	}
	return version
}
