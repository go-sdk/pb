package pb

var (
	VERSION = "latest"
	GITHASH = "b0af65c"
	BUILT   = "2022-01-28T02:14:51+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
