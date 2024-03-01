package codecov

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"os"
	"os/exec"
	"path"
)

var cliPkg, cliName string

func localUpload() error {
	if err := run("create-commit"); err != nil {
		return err
	}
	return nil
}

func run(args ...string) error {
	cli, err := bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cli, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func bin() (string, error) {
	return runfiles.Rlocation(path.Join(cliPkg, cliName))
}
