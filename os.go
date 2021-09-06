package os

import (
	"runtime"
)

// Info is used to store OS information
type Info struct {
	Name    string
	Version string
	Arch    string
}

func newInfo() Info {
	return Info{
		Name:    runtime.GOOS,
		Version: "unknown",
		Arch:    runtime.GOARCH,
	}
}

// GetInfo obtains operating system information
func GetInfo() (Info, error) {
	return getInfo()
}
