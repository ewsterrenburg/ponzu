package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/prometheus/common/log"
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
	gopath, err := getGOPATH()
	if err != nil {
		return err
	}
	repo := ponzuRepo

	local := filepath.Join(gopath, "src", filepath.Join(repo...))

	localMod := filepath.Join(gopath, "pkg/mod", filepath.Join(repo...))
	network := "https://" + strings.Join(repo, "/") + ".git"

	// create the directory or overwrite it
	err = os.MkdirAll(path, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	if dev {
		if fork != "" {
			local = filepath.Join(gopath, "src", fork)
		}

		err = execAndWait("git", "clone", local, "--branch", "ponzu-dev", "--single-branch", path)
		if err != nil {
			return err
		}

		err = vendorCorePackages(path)
		if err != nil {
			return err
		}

		fmt.Println("Dev build cloned from " + local + ":ponzu-dev")
		return nil
	}

	// try to git clone the repository from the local machine's $GOPATH
	err = execAndWait("git", "clone", local, path)
	if err != nil {

		fmt.Println("Couldn't clone from", local, "- trying modules...")
		err = execAndWait("git", "clone", localMod, path)
		if err != nil {

			fmt.Println("Couldn't clone from", local, "- trying network...", network)

			// try to git clone the repository over the network

			err = execAndWait("git", "clone", network, path)

			if err != nil {
				fmt.Println("Network clone failure.")
				// failed
				return fmt.Errorf("Failed to clone files from local machine [%s] and over the network [%s].\n%s", local, network, err)
			}
		}
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
		log.Fatal((err))
	}
	fmt.Println("creating mod file...")
	if err := execAndWait("go", "mod", "init", path); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New ponzu project created at", path)

	f, err := os.OpenFile("go.mod", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.WriteString("require github.com/padraicbc/ponzu/dynamic/content  v0.0.0\n" +
		"// Hack so we always use local generated on side effect import\n" +
		"replace github.com/padraicbc/ponzu/dynamic/content => ./dynamic/content"); err != nil {
		log.Fatal(err)
	}

	return nil
}

func init() {
	newCmd.Flags().StringVar(&fork, "fork", "", "modify repo source for Ponzu core development")
	newCmd.Flags().BoolVar(&dev, "dev", false, "modify environment for Ponzu core development")

	RegisterCmdlineCommand(newCmd)
}
