package main

import (
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	"io/fs"
	"path/filepath"
	"strings"
)

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
