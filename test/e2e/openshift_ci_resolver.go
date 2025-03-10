package e2e

import "knative.dev/kn-plugin-event/test/images"

func init() { //nolint:gochecknoinits
	resolver := &images.EnvironmentalBasedResolver{
		Prefix: "kn-plugin-event-test",
	}
	images.Resolvers = append(images.Resolvers, resolver)
}
