package codecov

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/davecgh/go-spew/spew"
	"os/exec"
	"path"
)

var cliPkg, cliName string

func run(args ...string) error {
	spew.Dump("hello")
	cli, err := codecovCliFromBazel()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(cli, args...)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	println(string(bytes))
	spew.Dump(err)
	spew.Dump("done")
	return nil
}

// codecovCli returns an absolute path to the codecov-cli tool
func codecovCli(pkg, name string) (string, error) {
	return runfiles.Rlocation(path.Join(pkg, name))
}

func codecovCliFromBazel() (string, error) {
	return codecovCli(cliPkg, cliName)
}
