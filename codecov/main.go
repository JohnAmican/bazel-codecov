package main

import (
	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/davecgh/go-spew/spew"
	"os/exec"
)

func main() {
	println("hello world")
	runfiles, err := bazel.ListRunfiles()
	if err != nil {
		panic(err)
	}
	var i int
	var cli string
	for _, runfile := range runfiles {
		if runfile.ShortPath == "rules_python_wheel_entry_point_codecovcli" {
			cli = runfile.Path
			i += 1
		}
	}
	if cli == "" {
		panic("no bin")
	}
	if i > 1 {
		panic("multiple bins")
	}
	cmd := exec.Command(cli, "--help")
	bytes, err := cmd.CombinedOutput()
	println(string(bytes))
	spew.Dump(err)
}
