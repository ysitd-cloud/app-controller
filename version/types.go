package version

import (
	"runtime"

	"k8s.io/apimachinery/pkg/version"
)

var major string
var minor string
var gitVersion string
var gitCommit string
var buildDate string

var Version version.Info = version.Info{
	Major:        major,
	Minor:        minor,
	GitVersion:   gitVersion,
	GitCommit:    gitCommit,
	GitTreeState: "clean",
	BuildDate:    buildDate,
	GoVersion:    runtime.Version(),
	Compiler:     runtime.Compiler,
	Platform:     runtime.GOOS,
}
