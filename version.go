package pb

var (
	VERSION = "latest"
	GITHASH = "b318bf4"
	BUILT   = "2022-01-28T02:18:55+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
