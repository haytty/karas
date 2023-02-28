package version

var (
	Name         = "karas"
	BuildVersion = ""
	Version      = "1.0.0"
)

func CurrentVersion() string {
	return Version
}
