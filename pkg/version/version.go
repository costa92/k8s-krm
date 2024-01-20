package version

var (
	Version = "dev"
)

// CheckNewVersion checks if a new version is available.
func CheckNewVersion() {
	if Version == "dev" {
		return
	}
}
