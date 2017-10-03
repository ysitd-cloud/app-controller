package version

import (
	"k8s.io/apimachinery/pkg/version"
	"runtime"
)

var Major string
var Minor string
var GitVersion string
var GitCommit string
var GitTreeState string = "clean"
var BuildDate string
var GoVersion string = runtime.Version()
var Compiler string = runtime.Compiler
var Platform string = runtime.GOOS

var Version version.Info = version.Info{
	Major:        Major,
	Minor:        Minor,
	GitVersion:   GitVersion,
	GitCommit:    GitCommit,
	GitTreeState: GitTreeState,
	BuildDate:    BuildDate,
	GoVersion:    GoVersion,
	Compiler:     Compiler,
	Platform:     Platform,
}
