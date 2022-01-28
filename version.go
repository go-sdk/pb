package pb

var (
	VERSION = "latest"
	GITHASH = "9a0b5ae"
	BUILT   = "2022-01-28T02:09:43+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
