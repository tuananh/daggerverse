package main

import (
	"context"
	"fmt"
)

type Apko struct{}

func (a *Apko) Build(
	ctx context.Context,
	source *Directory,
	// +default="amd64"
	arch string,
	apkoFile *File,
	// +optional
	// default="packages"
	packagesAppend *Directory,
	// +optional
	keyringAppend *File,
	image string,
	// +default="latest"
	tag string,
	// +optional
	tar string) *File {
	f, err := apkoFile.Name(ctx)
	if err != nil {
		panic(err)
	}

	cli := dag.Pipeline("apko-build")

	base := cli.Container().
		From("cgr.dev/chainguard/apko:latest")

	execOpts := []string{"build", fmt.Sprintf("%s", f), fmt.Sprintf("--arch=%s", arch)}

	if packagesAppend != nil {
		execOpts = append(execOpts, "-r", "./packages")
		base = base.WithMountedDirectory("/workspace/packages", packagesAppend)
	}

	if keyringAppend != nil {
		kr, err := keyringAppend.Name(ctx)
		if err != nil {
			panic(err)
		}
		execOpts = append(execOpts, "-k", fmt.Sprintf("%s", kr))
		base = base.WithMountedFile("/workspace/melange.rsa.pub", keyringAppend)
	}

	if image != "" {
		image = fmt.Sprintf("%s:%s", image, tag)
		execOpts = append(execOpts, image)
	}

	if tar == "" {
		tar = "apko.tar"
	}

	execOpts = append(execOpts, fmt.Sprintf("%s", tar))

	fmt.Println("execOpts", execOpts)

	return base.
		WithMountedDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec(execOpts).
		File(fmt.Sprintf("/workspace/%s", tar))

}
