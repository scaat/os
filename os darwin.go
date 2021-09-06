//go:build darwin
// +build darwin

package os

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func getInfo() (Info, error) {
	os := newInfo()
	cmd := exec.Command("sw_vers")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return os, fmt.Errorf("exec sw_vers error:%w", err)
	}

	versionSlice := regexp.MustCompile(`ProductVersion:(\d|.+)`).FindStringSubmatch(string(stdout))
	fmt.Println(versionSlice)
	if len(versionSlice) > 1 {
		os.Version = strings.TrimSpace(versionSlice[1])
	}
	return os, nil
}
