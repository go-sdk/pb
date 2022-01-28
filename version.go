package pb

var (
	VERSION = "latest"
	GITHASH = "b84b8c6"
	BUILT   = "2022-01-28T02:02:58+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
