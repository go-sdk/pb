cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "6727949"
	BUILT   = "2022-01-27T10:12:00+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
