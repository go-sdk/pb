package pb

var (
	VERSION = "latest"
	GITHASH = "50e87bc"
	BUILT   = "2022-01-28T01:38:13+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
