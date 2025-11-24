// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/x/internal/dagger"
)

type X struct{}

// The base development container.
func (m *X) devcontainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	// +optional
	platform *dagger.Platform,
) (*dagger.Container, error) {
	if platform == nil {
		enginePlatform, err := dag.DefaultPlatform(ctx)
		if err != nil {
			return nil, err
		}

		platform = &enginePlatform
	}

	return dag.Container(dagger.ContainerOpts{Platform: *platform}).
		From("golang:1.25-alpine").
		WithExec([]string{"apk", "add", "--no-cache", "git"}).
		WithMountedCache("/go/pkg/mod/", dag.CacheVolume("go-mod-125")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-125")).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithDirectory("/go/src/", source).
		WithWorkdir("/go/src/").
		WithExec([]string{"go", "mod", "download"}), nil
}

// TODO: digitalocean/gta
// TODO: caarlos0/svu
