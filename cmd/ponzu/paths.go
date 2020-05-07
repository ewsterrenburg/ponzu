package main

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// buildOutputName returns the correct ponzu-server file name
// based on the host Operating System
func buildOutputName() string {
	if runtime.GOOS == "windows" {
		return "ponzu-server.exe"
	}

	return "ponzu-server"
}

// resolve GOPATH. In 1.8 can be default, or custom. A custom GOPATH can
// also contain multiple paths, in which case 'go get' uses the first
func getGOPATH() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		// not set, find the default
		usr, err := user.Current()
		if err != nil {
			return gopath, err
		}
		return filepath.Join(usr.HomeDir, "go"), nil
	}
	// parse out in case of multiple, retain first
	if runtime.GOOS == "windows" {

		return strings.Split(gopath, ";")[0], nil
	}
	return strings.Split(gopath, ":")[0], nil

}
