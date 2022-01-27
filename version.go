cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "f2f92b2"
	BUILT   = "2022-01-27T09:42:18+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
