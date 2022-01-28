package pb

var (
	VERSION = "latest"
	GITHASH = "b1ec822"
	BUILT   = "2022-01-28T01:39:17+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
