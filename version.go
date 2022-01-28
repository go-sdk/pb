package pb

var (
	VERSION = "v0.2.0"
	GITHASH = "97360aa"
	BUILT   = "2022-01-28T16:20:03+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
