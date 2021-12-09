//go:build mage
// +build mage

package main

import (
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/artifact"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
	"github.com/wavesoftware/go-magetasks/pkg/checks"
	"github.com/wavesoftware/go-magetasks/pkg/knative"
	"knative.dev/kn-plugin-event/overrides"
	"knative.dev/kn-plugin-event/pkg/metadata"
)

// Default target is set to binary.
//goland:noinspection GoUnusedGlobalVariable
var Default = Build

// TODO: Remove tasks definitions in favor of mage:import comment
//       See: https://github.com/wavesoftware/go-magetasks/issues/29

// Clean will clean project files.
func Clean() {
	mg.Deps(magetasks.Clean)
}

// Generate will generate Go code.
func Generate() error {
	return sh.RunV("go", "generate", "./...")
}

// Check will run all lints checks.
func Check() {
	mg.Deps(magetasks.Check)
}

// Test will execute regular unit tests.
func Test() {
	mg.Deps(Generate, magetasks.Test)
}

// Build will build project artifacts, binaries and images.
func Build() {
	mg.Deps(Generate, magetasks.Build)
}

// Publish will publish built artifacts to remote site.
func Publish() {
	mg.Deps(magetasks.Publish)
}

func init() {
	sender := artifact.Image{
		Metadata: config.Metadata{Name: "kn-event-sender"},
		Architectures: []platform.Architecture{
			platform.AMD64, platform.ARM64, platform.S390X, platform.PPC64LE,
		},
	}
	cli := artifact.Binary{
		Metadata: config.Metadata{
			Name:           "kn-event",
			BuildVariables: cliBuildVariables(sender),
		},
		Platforms: []artifact.Platform{
			{OS: platform.Linux, Architecture: platform.AMD64},
			{OS: platform.Linux, Architecture: platform.ARM64},
			{OS: platform.Linux, Architecture: platform.PPC64LE},
			{OS: platform.Linux, Architecture: platform.S390X},
			{OS: platform.Mac, Architecture: platform.AMD64},
			{OS: platform.Mac, Architecture: platform.ARM64},
			{OS: platform.Windows, Architecture: platform.AMD64},
		},
	}
	magetasks.Configure(config.Config{
		Version: &config.Version{
			Path:     metadata.VersionPath(),
			Resolver: knative.NewVersionResolver(),
		},
		Artifacts: []config.Artifact{sender, cli},
		Checks:    []config.Task{checks.GolangCiLint()},
		BuildVariables: map[string]config.Resolver{
			metadata.ImageBasenamePath(): imageBasenameFromEnv,
		},
		Overrides: overrides.List,
	})
}

func cliBuildVariables(sender artifact.Image) config.BuildVariables {
	return config.NewBuildVariablesBuilder().
		ConditionallyAdd(
			referenceImageByDigest,
			metadata.ImagePath(),
			artifact.ImageReferenceOf(sender),
		).Build()
}

func imageBasenameFromEnv() string {
	return env("KO_DOCKER_REPO", "IMAGE_BASENAME")
}

func env(keys ...string) string {
	for _, key := range keys {
		if val, ok := os.LookupEnv(key); ok {
			return val
		}
	}
	return ""
}

func skipImageReference() bool {
	if val, ok := os.LookupEnv("SKIP_IMAGE_REFERENCE"); ok {
		return strings.ToLower(val) == "true"
	}
	return false
}

func referenceImageByDigest() bool {
	return !skipImageReference()
}
