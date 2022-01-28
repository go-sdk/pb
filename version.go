package pb

var (
	VERSION = "latest"
	GITHASH = "1571ad8"
	BUILT   = "2022-01-28T01:58:46+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
