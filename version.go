package pb

var (
	VERSION = "latest"
	GITHASH = "da72d4d"
	BUILT   = "2022-01-28T01:51:36+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
