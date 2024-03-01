package main

import (
	"context"
	"fmt"
)

type Grype struct{}

func (t *Grype) Base(
	// +optional
	// +default="latest"
	imageTag string,
) *Container {
	return dag.Container().
		From(fmt.Sprintf("anchore/grype:%s", imageTag))
}

func (t *Grype) Scan(
	ctx context.Context,
	imageRef string,
	// +default="latest"
	imageTag string,
) (string, error) {
	return t.Base(imageTag).
		WithExec([]string{imageRef}).
		Stdout(ctx)
}
