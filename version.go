package pb

var (
	VERSION = "latest"
	GITHASH = "c7d462e"
	BUILT   = "2022-01-28T01:52:24+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
