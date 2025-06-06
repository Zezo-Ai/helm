/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ensure

import (
	"os"
	"path/filepath"
	"testing"

	"helm.sh/helm/v4/pkg/helmpath"
	"helm.sh/helm/v4/pkg/helmpath/xdg"
)

// HelmHome sets up a Helm Home in a temp dir.
func HelmHome(t *testing.T) {
	t.Helper()
	base := t.TempDir()
	t.Setenv(xdg.CacheHomeEnvVar, base)
	t.Setenv(xdg.ConfigHomeEnvVar, base)
	t.Setenv(xdg.DataHomeEnvVar, base)
	t.Setenv(helmpath.CacheHomeEnvVar, "")
	t.Setenv(helmpath.ConfigHomeEnvVar, "")
	t.Setenv(helmpath.DataHomeEnvVar, "")
}

// TempFile ensures a temp file for unit testing purposes.
//
// It returns the path to the directory (to which you will still need to join the filename)
//
// The returned directory is automatically removed when the test and all its subtests complete.
//
//	tempdir := TempFile(t, "foo", []byte("bar"))
//	filename := filepath.Join(tempdir, "foo")
func TempFile(t *testing.T, name string, data []byte) string {
	t.Helper()
	path := t.TempDir()
	filename := filepath.Join(path, name)
	if err := os.WriteFile(filename, data, 0o755); err != nil {
		t.Fatal(err)
	}
	return path
}
