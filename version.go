package pb

var (
	VERSION = "latest"
	GITHASH = "7f65364"
	BUILT   = "2022-01-28T02:08:39+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
