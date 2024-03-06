package main

import (
	"context"
	"fmt"
)

type Apko struct{}

func (a *Apko) Build(
	ctx context.Context,
	source *Directory,
	apkoFile *File,
	// +optional
	// default="packages"
	packageAppend *Directory,
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
	base := dag.Container().
		From("cgr.dev/chainguard/apko:latest")

	execOpts := []string{"build", fmt.Sprintf("%s", f)}

	if packageAppend != nil {
		// execOpts = append(execOpts, "-p", fmt.Sprintf("%s", packageAppend))
		base = base.WithMountedDirectory("/packages", packageAppend)
	}

	if keyringAppend != nil {
		kr, err := keyringAppend.Name(ctx)
		if err != nil {
			panic(err)
		}
		execOpts = append(execOpts, "-k", fmt.Sprintf("%s", kr))
		base = base.WithMountedFile("/workspace", keyringAppend)
	}

	if image != "" {
		image = fmt.Sprintf("%s:%s", image, tag)
		execOpts = append(execOpts, image)
	}

	if tar == "" {
		tar = "apko.tar"
	}

	execOpts = append(execOpts, fmt.Sprintf("%s", tar))

	return base.
		WithMountedDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec(execOpts).
		File(fmt.Sprintf("/workspace/%s", tar))

}
