package pb

var (
	VERSION = "latest"
	GITHASH = "2d17b1c"
	BUILT   = "2022-01-28T01:57:59+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
