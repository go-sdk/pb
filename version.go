package pb

var (
	VERSION = "v0.1.5"
	GITHASH = "ca65418"
	BUILT   = "2022-01-28T06:19:36+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
