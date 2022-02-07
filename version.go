package pb

var (
	VERSION = "dev"
	GITHASH = "06847db"
	BUILT   = "2022-02-07T14:25:23+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
