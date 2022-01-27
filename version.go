cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "eb16c78"
	BUILT   = "2022-01-27T10:10:10+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
