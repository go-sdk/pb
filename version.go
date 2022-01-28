package pb

var (
	VERSION = "dev"
	GITHASH = "d6bd188"
	BUILT   = "2022-01-28T16:15:23+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
