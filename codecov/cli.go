package codecov

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"os"
	"os/exec"
	"strings"
)

// cliEntryPoint is defined in the go_library rule for this file using @pypi//:requirements.bzl%entry_point
var cliEntryPoint string = "__DEFINE_IN_X_DEFS__"

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
	//"cliPkg": "pypi_codecov_cli",
	//"cliName": "rules_python_wheel_entry_point_codecovcli",
	//x := strings.ReplaceAll(strings.TrimPrefix(cliEntryPoint, "@"), "//:", "/")
	//if x != "pypi_codecov_cli/rules_python_wheel_entry_point_codecovcli"
	return runfiles.Rlocation(strings.ReplaceAll(strings.TrimPrefix(cliEntryPoint, "@"), "//:", "/"))
}
