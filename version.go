package pb

var (
	VERSION = "dev"
	GITHASH = "3304ec8"
	BUILT   = "2022-01-28T06:23:30+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
