cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "5fb1d4f"
	BUILT   = "2022-01-27T09:52:45+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
