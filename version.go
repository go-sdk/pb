cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "3ef0405"
	BUILT   = "2022-01-27T10:02:36+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
