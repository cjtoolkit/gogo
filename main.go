package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func seekGogoFile() string {
	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gnode := filepath.FromSlash("/.gogo")
	lastPath := false
	for {
		if _, err := os.Stat(curdir + gnode); os.IsNotExist(err) {
			curdir = filepath.Dir(curdir)
			if lastPath || curdir == "." {
				log.Fatal("Could not find '.gogo'")
			} else if strings.Trim(curdir, "/") == "" || (runtime.GOOS == "windows" && len(strings.Trim(curdir, "\\")) == 2) {
				lastPath = true
			}
			continue
		}
		break
	}
	return curdir + gnode
}

func main() {
	gogoFile := seekGogoFile()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(gogoFile)
	if err != nil {
		log.Fatal(err)
	}
	cmdName := func() string {
		b, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		return strings.TrimSpace(string(b))
	}()
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	sdkPath := homeDir + filepath.FromSlash("/sdk/" +cmdName)
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		setupSdk(cmdName, homeDir)
	}

	cmdRun(cmdName, os.Args[1:]...)
}

func setupSdk(cmdName, homeDir string) {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// To work round go.mod
	defer os.Chdir(curDir)
	err = os.Chdir(homeDir)
	if err != nil {
		log.Fatal(err)
	}
	cmdGoGet(cmdName)
	cmdRun(cmdName, "download")
}

func cmdRun(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func cmdGoGet(cmdName string) {
	cmd := exec.Command("go", "get", "golang.org/dl/" + cmdName)
	cmd.Env = append(os.Environ(), "GOOS="+runtime.GOOS, "GOARCH="+runtime.GOARCH)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}