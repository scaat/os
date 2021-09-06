//go:build linux
// +build linux

package os

import (
	"fmt"
	"os/exec"
	"regexp"
)

func getInfo() (OSInfo, error) {
	os := newInfo()

	cmd := exec.Command("cat", "/etc/os-release")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return os, fmt.Errorf("exec cat error:%w", err)
	}

	id := regexp.MustCompile(`\nID="?(.*?)"?\n`).FindStringSubmatch(string(stdout))
	if len(id) > 1 {
		os.Name = id[1]
	}

	version := regexp.MustCompile(`VERSION_ID="?([.0-9]+)"?\n`).FindStringSubmatch(string(stdout))
	if len(version) > 1 {
		os.Version = version[1]
	}

	return os, nil
}
