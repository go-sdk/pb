package pb

var (
	VERSION = "latest"
	GITHASH = "f9c7b68"
	BUILT   = "2022-01-28T02:13:57+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
