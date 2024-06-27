package main

import (
	"context"
	"fmt"
	"path/filepath"
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
	// generate public/private key pair
	cli := dag.Pipeline("melange-build")

	ctr := cli.Container().From(fmt.Sprintf("cgr.dev/chainguard/melange:%s", imageTag)).
		WithWorkdir("/workspace").
		WithExec([]string{
			"keygen"})

	f, _ := melangeFile.Name(ctx)

	c := cli.Container().
		From(fmt.Sprintf("cgr.dev/chainguard/melange:%s", imageTag)).
		WithMountedDirectory("/workspace", workspaceDir).
		WithDirectory("/workspace", ctr.Directory("/workspace")).
		WithWorkdir("/workspace").
		WithExec([]string{
			"build", fmt.Sprintf("%s", f), "--arch", arch, "--signing-key=melange.rsa"},
			ContainerWithExecOpts{
				ExperimentalPrivilegedNesting: true,
				InsecureRootCapabilities:      true,
			})

	pk := c.File(filepath.Join("/workspace", "melange.rsa.pub"))

	return c.Directory("/workspace").WithFile(".", pk)
}
