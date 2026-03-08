package version

// Injected at build time via ldflags. See Makefile / .goreleaser.yml.
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

// ImageTag returns the Docker image tag corresponding to this CLI version.
// Release builds (e.g. "0.1.0") use the version directly; dev builds fall
// back to "latest".
func ImageTag() string {
	if Version == "dev" {
		return "latest"
	}
	return Version
}
