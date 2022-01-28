package pb

var (
	VERSION = "latest"
	GITHASH = "dc58538"
	BUILT   = "2022-01-28T01:45:55+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
