package metadata

import (
	"fmt"

	"github.com/blang/semver/v4"
)

// floatDirection defines the way to float a non-release version.
type floatDirection int

const (
	// floatDirectionUp means the minor version will be incremented to find
	// compatible range ,effectively meaning a next minor release.
	floatDirectionUp floatDirection = iota
	// floatDirectionDown means the minor version will be left intact, but patch
	// number will be removed, effectively meaning latest version from current
	// minor release.
	floatDirectionDown
)

// floatToRelease will build a full image name from basename, name, separator,
// version parts given as arguments. If version is a non-release it will be
// floated either up or down depending on direction argument. Floating up means
// to increase the minor number by 1. Floating down means leaving minor number
// as it was.
func floatToRelease(basename, name, separator, version string, direction floatDirection) string {
	sver, err := semver.ParseTolerant(version)
	if err == nil {
		version = fmt.Sprintf("v%d.%d.%d", sver.Major, sver.Minor, sver.Patch)
		if len(sver.Pre) > 0 || len(sver.Build) > 0 {
			// non release image
			major := sver.Major
			minor := sver.Minor
			if direction == floatDirectionUp {
				minor++
			}
			version = fmt.Sprintf("v%d.%d", major, minor)
		}
	}
	return fmt.Sprintf("%s%s%s:%s", basename, separator, name, version)
}
