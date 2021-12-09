package metadata

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// Version holds application version information.
var Version = "0.0.0" //nolint:gochecknoglobals

// VersionPath return a path to the version variable.
func VersionPath() string {
	return importPath("Version")
}

// ResolveVersion will try to lookup a kn-plugin-event version in case it's
// embedded in kn-cli.
func ResolveVersion() string {
	info, _ := debug.ReadBuildInfo()
	if info != nil {
		for _, dep := range info.Deps {
			if strings.Contains(dep.Path, "knative.dev/kn-plugin-event") {
				return versionOf(dep)
			}
		}
	}
	return Version
}

func versionOf(dep *debug.Module) string {
	if dep.Replace == nil {
		return dep.Version
	}
	return fmt.Sprintf("%s-%s", dep.Replace.Version, dep.Replace.Sum)
}
