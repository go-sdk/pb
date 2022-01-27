cat >pb/version.go <<EOF
package pb

var (
	VERSION = "latest"
	GITHASH = "e67f938"
	BUILT   = "2022-01-27T09:58:30+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
