package pb

var (
	VERSION = "latest"
	GITHASH = "8f0f25a"
	BUILT   = "2022-01-28T02:02:16+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
