package golang

import (
	"eol-checker/internal/api"
	"eol-checker/internal/config"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const productEndpoint = "golang"

type GolangChecker struct {
}

func getVersionFromFile(path string) (string, error) {
	// this should pull from the config
	// or the config should be parsed on running and versions passed as param to each product
	// use this for now

	// read golang version from go.mod file
	// this can also probably be checked using a go.version command
	// maybe have config get to choose what way

	// TODO: make sure we have path if this is chosen

	file, err := os.ReadFile(path)
	if err != nil {
		slog.Error("error reading go.mod file", "error", err)
		return "", err
	}

	fileLines := strings.Split(string(file), "\n")
	for _, line := range fileLines {
		if len(line) > 2 {
			if line[0:2] == "go" { // TODO: make this safer
				strVer := strings.Join(strings.Split(strings.Split(line, " ")[1], ".")[0:2], ".")
				// ver, err := strconv.ParseFloat(strVer, 64)
				// if err != nil {
				// 	slog.Error(fmt.Sprintf("error passing go version %v into float64", strVer), "error", err)
				// }
				return strVer, nil
			}
		}
	}

	err = fmt.Errorf("go version no found in go.mod")
	slog.Error("error getting go version", "error", err)

	return "", err
}

func getVersionFromCommand() (string, error) {
	res := exec.Command("go", "version")
	if res.Err != nil {
		slog.Error("error calling `go version`", "error", res.Err)
		return "", res.Err
	}
	out, err := res.Output()
	if err != nil {
		slog.Error("error getting `go version` output", "error", err)
		return "", err
	}

	versionRegex, err := regexp.Compile(`([0-9]+\.[0-9]+)`)
	if err != nil {
		slog.Error("failed to compile regex check", "error", err)
		return "", err
	}
	versionString := versionRegex.Find(out)
	if versionString == nil {
		slog.Error("failed to find a version in command output")
		return "", fmt.Errorf("failed to find a version in command output")
	}

	return string(versionString), nil
}

func (g GolangChecker) IsEol(c config.ProductConfig) (*api.EolCheck, error) {
	var version string
	if c.Method == "file" {
		v, err := getVersionFromFile(c.Path)
		if err != nil {
			return nil, err
		}
		version = v
	} else { // else if check=command
		v, err := getVersionFromCommand()
		if err != nil {
			return nil, err
		}
		version = v
	}

	versionData, err := api.DoEolCheck(productEndpoint, version)
	if err != nil {
		return nil, err
	}

	versionData.CheckId = c.Id
	versionData.Product = c.Product

	return versionData, nil
}
