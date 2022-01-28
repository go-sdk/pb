package pb

var (
	VERSION = "dev"
	GITHASH = "916f7e7"
	BUILT   = "2022-01-28T02:40:13+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
