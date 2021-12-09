//go:generate go run knative.dev/kn-plugin-event/tools/builttime

package metadata

var (
	// Image holds information about companion image reference.
	Image = "" //nolint:gochecknoglobals
	// ImageBasename holds a basename of a image, so the development reference
	// could be built from it.
	ImageBasename = "" //nolint:gochecknoglobals
	// ImageBasenameSeparator holds a separator between image basename and name.
	// nolint:gochecknoglobals
	ImageBasenameSeparator = "/"
)

// ResolveImage will try to resolve the image reference from set values. If
// Image is given it will be used, otherwise the ImageBasename and Version will
// be used.
func ResolveImage() string {
	image := Image
	//goland:noinspection GoBoolExpressions
	if imageFromEnvironmentSet {
		image = imageFromEnvironment
	}
	if image == "" {
		name := "kn-event-sender"
		return floatToRelease(ImageBasename, name, ImageBasenameSeparator,
			ResolveVersion(), floatDirectionDown)
	}
	return image
}

// ImageOverriden check if image has been overridden on build time.
func ImageOverriden() bool {
	//goland:noinspection GoBoolExpressions
	return imageFromEnvironmentSet
}

// ImagePath return a path to the image variable.
func ImagePath() string {
	return importPath("Image")
}

// ImageBasenamePath return a path to the image basename variable.
func ImageBasenamePath() string {
	return importPath("ImageBasename")
}
