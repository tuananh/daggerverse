package main

import (
	"context"
	"fmt"
)

type Melange struct{}

func (m *Melange) Build(
	ctx context.Context,
	melangeFile *File,
	workspaceDir *Directory,
	// +default="amd64"
	arch string,
	// +default="latest"
	imageTag string,
) *Directory {
	return dag.Container().
		From(fmt.Sprintf("cgr.dev/chainguard/melange:%s", imageTag)).
		WithMountedDirectory("/workspace", workspaceDir).
		WithMountedFile("/workspace/melange.yaml", melangeFile).
		WithWorkdir("/workspace").
		WithExec([]string{
			"build", "/workspace/melange.yaml", "--arch", arch},
			ContainerWithExecOpts{
				ExperimentalPrivilegedNesting: true,
				InsecureRootCapabilities:      true,
			}).
		Directory("/workspace/packages")
}
