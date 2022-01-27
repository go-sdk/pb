cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "fdc5c13"
	BUILT   = "2022-01-27T10:01:09+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
