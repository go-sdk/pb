package pb

var (
	VERSION = "v0.1.4"
	GITHASH = "4d81fa5"
	BUILT   = "2022-01-28T03:11:42+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
