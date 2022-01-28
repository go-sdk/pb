package pb

var (
	VERSION = "dev"
	GITHASH = "7dd60bd"
	BUILT   = "2022-01-28T16:18:16+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
