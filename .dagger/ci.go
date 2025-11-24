// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/x/internal/dagger"
)

// Run all tests. Defaults to the repository base directory.
func (m *X) Test(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	ctr, err := m.devcontainer(ctx, source, nil)
	if err != nil {
		return "", err
	}

	return ctr.WithExec([]string{"go", "test", "-cover", "./..."}).CombinedOutput(ctx)
}
