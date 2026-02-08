package golang_methods

import (
	"errors"
	"log/slog"
	"os"
	"strings"
)

type File struct{}

// product: golang
// method: file
// args:
//	0 - path STRING

func parseArgs(args []any) (string, error) {
	var path string
	var err error
	if len(args) != 1 {
		return "", errors.New("not enough arguments passed into function")
	}
	path, ok := args[0].(string)
	if !ok {
		return "", errors.New("argument passed as PATH cannot be cast to string")
	}
	return path, err
}

func (m File) Run(id string, args []any) (string, error) {
	logger := slog.Default().With("product", "golang", "method", "file", "check_id", id)
	path, err := parseArgs(args)
	if err != nil {
		logger.Error("failed to parse argments", "error", err)
		return "", err
	}

	path, ok := args[0].(string)
	if !ok {
		errorMsg := "failed to assert path argument to string"
		logger.Error(errorMsg)
		return "", errors.New(errorMsg)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		logger.Error("error reading go.mod file", "path", path, "error", err)
		return "", err
	}

	fileLines := strings.Split(string(file), "\n")
	for _, line := range fileLines {
		if len(line) > 2 {
			if line[0:2] == "go" { // TODO: make this safer, regex check?
				strVer := strings.Join(strings.Split(strings.Split(line, " ")[1], ".")[0:2], ".")
				return strVer, nil
			}
		}
	}

	// if not returned in loop, failed to find version
	errMsg := "failed to find go version in go.mod file"
	logger.Error(errMsg)

	return "", err
}
