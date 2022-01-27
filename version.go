cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "3d3c412"
	BUILT   = "2022-01-27T09:46:58+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
