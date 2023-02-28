package version

var (
	Name         = "golang_cli_template"
	BuildVersion = ""
	Version      = "1.0.0"
)

func CurrentVersion() string {
	return Version
}
