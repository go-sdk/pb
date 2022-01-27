cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "a5b6f95"
	BUILT   = "2022-01-27T09:56:22+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
