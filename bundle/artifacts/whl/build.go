package whl

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/databricks/cli/bundle"
	"github.com/databricks/cli/bundle/config"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/diag"
	"github.com/databricks/cli/libs/log"
	"github.com/databricks/cli/libs/python"
)

type build struct {
	name string
}

func Build(name string) bundle.Mutator {
	return &build{
		name: name,
	}
}

func (m *build) Name() string {
	return fmt.Sprintf("artifacts.whl.Build(%s)", m.name)
}

func (m *build) Apply(ctx context.Context, b *bundle.Bundle) diag.Diagnostics {
	artifact, ok := b.Config.Artifacts[m.name]
	if !ok {
		return diag.Errorf("artifact doesn't exist: %s", m.name)
	}

	cmdio.LogString(ctx, fmt.Sprintf("Building %s...", m.name))

	out, err := artifact.Build(ctx)
	if err != nil {
		return diag.Errorf("build failed %s, error: %v, output: %s", m.name, err, out)
	}
	log.Infof(ctx, "Build succeeded")

	dir := artifact.Path
	distPath := filepath.Join(artifact.Path, "dist")
	wheels := python.FindFilesWithSuffixInPath(distPath, ".whl")
	if len(wheels) == 0 {
		return diag.Errorf("cannot find built wheel in %s for package %s", dir, m.name)
	}
	for _, wheel := range wheels {
		artifact.Files = append(artifact.Files, config.ArtifactFile{
			Source: wheel,
		})
	}

	return nil
}
