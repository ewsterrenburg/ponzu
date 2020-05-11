package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [flags] <project name>",
	Short: "creates a project directory of the name supplied as a parameter",
	Long: `Creates a project directory of the name supplied as a parameter
immediately following the 'new' option. Note:
'new' depends on the program 'git' and possibly a network connection. If
there is no local repository to clone from at the local machine's $GOPATH,
'new' will attempt to clone the 'github.com/padraicbc/ponzu' package from
over the network.`,
	Example: `$ ponzu new github.com/nilslice/proj
> New ponzu project created at github.com/nilslice/proj`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := "ponzu"
		if len(args) > 0 {
			projectName = args[0]
		} else {
			msg := "Please provide a project name."
			msg += "\nThis will create a directory within your pwd."
			return fmt.Errorf("%s", msg)
		}
		return newProjectInDir(projectName)
	},
}

func newProjectInDir(path string) error {
	_, err := os.Stat(path)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// path exists, ask if it should be overwritten
	if err == nil {
		fmt.Printf("Using '%s' as project directory\n", path)
		fmt.Println("Path exists, overwrite contents? (y/N):")
		answer, err := getAnswer()
		if err != nil {
			return err
		}

		switch answer {
		case "n", "no", "\r\n", "\n", "":
			fmt.Println("")

		case "y", "yes":
			err := os.RemoveAll(path)
			if err != nil {
				return fmt.Errorf("Failed to overwrite %s. \n%s", path, err)
			}

			return createProjectInDir(path)

		default:
			fmt.Println("Input not recognized. No files overwritten. Answer as 'y' or 'n' only.")
		}

		return nil
	}
	return createProjectInDir(path)
}

func createProjectInDir(path string) error {

	repo := ponzuRepo

	network := "https://" + strings.Join(repo, "/") + ".git"

	// create the directory or overwrite it
	err := os.MkdirAll(path, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	// try to git clone the repository over the network

	err = execAndWait("git", "clone", network, path)

	if err != nil {
		fmt.Println("Network clone failure.")
		// failed
		return fmt.Errorf("Failed to clone files from  network [%s].\n%s", network, err)
	}

	// create an directory in ./cmd/ponzu and move content,
	// management and system packages into it
	err = vendorCorePackages(path)
	if err != nil {
		return err
	}

	// remove non-project files and directories
	rmPaths := []string{".git", ".circleci", "go.mod", "go.sum"}
	for _, rm := range rmPaths {

		dir := filepath.Join(path, rm)

		err = os.RemoveAll(dir)
		if err != nil {
			fmt.Println("Failed to remove directory from your project path. Consider removing it manually:", dir)
		}
	}
	// change dir and create mod file
	if err := os.Chdir(path); err != nil {
		return err
	}
	fmt.Println("creating mod file...")
	if err := execAndWait("go", "mod", "init", path); err != nil {
		return err
	}
	fmt.Println("New ponzu project created at", path)

	f, err := os.OpenFile("go.mod", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.WriteString("require github.com/padraicbc/ponzu/content  v0.0.0\n" +
		"// Hack so we always use local generated on side effect import\n" +
		"replace github.com/padraicbc/ponzu/dynamic/content => ./content"); err != nil {
		return err
	}

	return nil
}

func init() {
	newCmd.Flags().StringVar(&fork, "fork", "", "modify repo source for Ponzu core development")
	newCmd.Flags().BoolVar(&dev, "dev", false, "modify environment for Ponzu core development")

	RegisterCmdlineCommand(newCmd)
}
