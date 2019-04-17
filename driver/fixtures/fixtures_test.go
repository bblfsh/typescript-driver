package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/sdk/v3/driver"
	"github.com/bblfsh/sdk/v3/driver/fixtures"
	"github.com/bblfsh/sdk/v3/driver/native"
	"github.com/bblfsh/typescript-driver/driver/normalizer"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "typescript",
	Ext:  ".ts",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return native.NewDriverAt(filepath.Join(projectRoot, "build/bin/native"), native.UTF8)
	},
	Transforms: normalizer.Transforms,
	BenchName:  "types",
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			// TODO: list native types that should be converted to semantic UAST
		},
	},
}

func TestTypescriptDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkTypescriptDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
