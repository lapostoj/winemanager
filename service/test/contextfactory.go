package test

import (
	"context"

	"google.golang.org/appengine/aetest"
)

// AnAppEngineTestContext returns an aetest Context to use in tests
func AnAppEngineTestContext() context.Context {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		panic(err)
	}
	defer done()
	return ctx
}
