package codecov

import (
	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"testing"
)

func TestRun(t *testing.T) {
	dir, err := bazel.NewTmpDir("whocares")
	if err != nil {
		t.Fatal(err)
	}
	println("okokokoko")
	println(dir)
	println("okokokoko")
	if err = run(
		"--help",
		"--verbose",
		"upload-process",
		"--fail-on-error",
		"-C", "c1d4af790b8d735224fc7f190cbbdcc99a71462c",
		"--disable-search",
		"-f", "herpderp.dat",
	); err != nil {
		t.Fatal(err)
	}
}
