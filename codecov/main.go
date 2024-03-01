package main

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/davecgh/go-spew/spew"
	"os/exec"
	"path"
)

var cliPkg, cliName string

func main() {
	printRunfilesEnv()
	do()
}

func do() {
	cli, err := codecovCliFromBazel()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(cli, "--help")
	bytes, err := cmd.CombinedOutput()
	println(string(bytes))
	spew.Dump(err)
}

// codecovCli returns an absolute path to the codecov-cli tool
func codecovCli(pkg, name string) (string, error) {
	return runfiles.Rlocation(path.Join(pkg, name))
}

func codecovCliFromBazel() (string, error) {
	return codecovCli(cliPkg, cliName)
}
