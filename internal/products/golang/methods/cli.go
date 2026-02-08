package golang_methods

import (
	"fmt"
	"log/slog"
	"os/exec"
	"regexp"
)

type CLI struct{}

func (c CLI) Run(id string, args []any) (string, error) {
	logger := slog.Default().With("product", "golang", "method", "command", "check_id", id)
	res := exec.Command("go", "version")
	if res.Err != nil {
		logger.Error("error calling `go version`", "error", res.Err)
		return "", res.Err
	}
	out, err := res.Output()
	if err != nil {
		logger.Error("error getting `go version` output", "error", err)
		return "", err
	}

	versionRegex, err := regexp.Compile(`([0-9]+\.[0-9]+)`)
	if err != nil {
		logger.Error("failed to compile regex check", "error", err)
		return "", err
	}
	versionString := versionRegex.Find(out)
	if versionString == nil {
		logger.Error("failed to find a version in command output")
		return "", fmt.Errorf("failed to find a version in command output")
	}

	return string(versionString), nil
}
