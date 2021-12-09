package metadata_test

import (
	"fmt"
	"sync"
	"testing"

	"gotest.tools/v3/assert"
	"knative.dev/kn-plugin-event/pkg/metadata"
)

const exampleRegistryPath = "oci.knative.dev/kn-plugin-event"

var mx = sync.Mutex{} //nolint:gochecknoglobals

func TestResolveImage(t *testing.T) {
	for _, testCase := range testResolveImageCases() {
		tc := testCase
		name := fmt.Sprintf("%v:%v:%v", tc.version, tc.image, tc.imageBasename)
		t.Run(name, func(t *testing.T) {
			if metadata.ImageOverriden() {
				t.Skip("Image is overridden")
			}
			withTestValues(tc.testMetadata, func() {
				got := metadata.ResolveImage()
				assert.Equal(t, got, tc.want)
			})
		})
	}
}

func testResolveImageCases() []testResolveImageCase {
	return []testResolveImageCase{{
		testMetadata{
			version:       "v1.2.3-12-g456fb21",
			imageBasename: exampleRegistryPath,
		},
		"oci.knative.dev/kn-plugin-event/kn-event-sender:v1.2",
	}, {
		testMetadata{
			version:       "v1.2.3",
			imageBasename: exampleRegistryPath,
		},
		"oci.knative.dev/kn-plugin-event/kn-event-sender:v1.2.3",
	}, {
		testMetadata{
			image: "image.registry.local@sha256:c1c34d43225f245f154a4db9f7e5f9c1dd6cb7ad63e010cdfa960fab6bdefd44",
		},
		"image.registry.local@sha256:c1c34d43225f245f154a4db9f7e5f9c1dd6cb7ad63e010cdfa960fab6bdefd44",
	}}
}

type testResolveImageCase struct {
	testMetadata
	want string
}

type testMetadata struct {
	version       string
	image         string
	imageBasename string
}

func withTestValues(vals testMetadata, fn func()) {
	mx.Lock()
	defer mx.Unlock()
	old := testMetadata{
		version:       metadata.Version,
		image:         metadata.Image,
		imageBasename: metadata.ImageBasename,
	}
	defer func() {
		metadata.Version = old.version
		metadata.Image = old.image
		metadata.ImageBasename = old.imageBasename
	}()
	metadata.Version = vals.version
	metadata.Image = vals.image
	metadata.ImageBasename = vals.imageBasename
	fn()
}
