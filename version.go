package pb

var (
	VERSION = "latest"
	GITHASH = "a53db11"
	BUILT   = "2022-01-28T01:50:44+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
