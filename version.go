package pb

var (
	VERSION = "dev"
	GITHASH = "ddf8082"
	BUILT   = "2022-01-28T06:18:22+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
