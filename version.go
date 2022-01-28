package pb

var (
	VERSION = "latest"
	GITHASH = "5dcc2c3"
	BUILT   = "2022-01-28T01:45:06+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
