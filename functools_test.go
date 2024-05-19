package functools_test

import (
	"os"
	"runtime"
	"testing"
)

func Test_RuntimeVersion(t *testing.T) {
	isCiRun := os.Getenv("CI") != ""
	if !isCiRun {
		t.Skip("Skipping test because it is not running in CI environment.")
	}
	versionInfo := runtime.Version()
	t.Logf("Go version: %s", versionInfo)
}
