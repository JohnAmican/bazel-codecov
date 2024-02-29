package main

import (
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/davecgh/go-spew/spew"
	"io/fs"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var cliPkg, cliName string

func main() {
	printRunfilesEnv()
	do()
}

func printRunfilesEnv() {
	kvs, err := runfiles.Env()
	if err != nil {
		panic(err)
	}
	for _, kv := range kvs {
		if k, v, found := strings.Cut(kv, "="); found {
			switch k {
			case "RUNFILES_DIR":
				println("runfiles:")
				err = filepath.WalkDir(v, func(path string, d fs.DirEntry, err error) error {
					//if strings.HasSuffix(path, "coverage.dat") {
					if strings.Contains(path, ".dat") {
						println(path)
					}
					//}
					return nil
				})
			default:
				println(fmt.Sprintf("%s=%s", k, v))
			}
		}
	}
	if err != nil {
		panic(err)
	}
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
