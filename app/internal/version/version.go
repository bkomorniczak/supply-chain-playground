package version

import "runtime"

var (
	commit = "dev"
	buildTime = "unknown"
)

type V struct {
	Commit string
	BuildTime string
	GoVersion string
}

func Info() V {
	return V{
		Commit:    commit,
		BuildTime: buildTime,
		GoVersion: runtime.Version(),
	}
}

func (v V) String() string {
	return v.Commit + "@" + v.BuildTime
}

func (v V) JSON() map[string]string {
	return map[string]string{
		"commit":     v.Commit,
		"build_time": v.BuildTime,
		"go_version": v.GoVersion,
	}
}