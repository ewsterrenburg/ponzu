package main

import (
	"strings"

	"github.com/spf13/cobra"
)

func buildPonzuServer() error {
	// copy all ./content files to internal directory

	// copy all ./addons files & dirs to internal  directory
	src := "addons"
	dst := "vendor"

	err := copyFilesWarnConflicts(src, dst, nil)
	if err != nil {
		return err
	}

	// execute go build -o ponzu cmd/ponzu/*.go
	cmdPackageName := strings.Join([]string{".", "cmd", "ponzu"}, "/")
	buildOptions := []string{"build", "-o", buildOutputName(), cmdPackageName}
	return execAndWait(gocmd, buildOptions...)
}

var buildCmd = &cobra.Command{
	Use:   "build [flags]",
	Short: "build will build/compile the project to then be run.",
	Long: `From within your Ponzu project directory, running build will copy and move
the necessary files from your workspace into the vendored directory, and
will build/compile the project to then be run.

By providing the 'gocmd' flag, you can specify which Go command to build the
project, if testing a different release of Go.

Errors will be reported, but successful build commands return nothing.`,
	Example: `$ ponzu build
(or)
$ ponzu build --gocmd=go1.8rc1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildPonzuServer()
	},
}

func init() {
	RegisterCmdlineCommand(buildCmd)
}
