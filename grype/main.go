package main

import (
	"context"
	"fmt"
)

type Grype struct{}

func (g *Grype) Scan(
	ctx context.Context,
	imageRef string,
	// +default="latest"
	imageTag string,
) (string, error) {
	return dag.Container().
		From(fmt.Sprintf("anchore/grype:%s", imageTag)).
		WithExec([]string{imageRef}).
		Stdout(ctx)
}
