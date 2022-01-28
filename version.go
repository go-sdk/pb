package pb

var (
	VERSION = "v0.1.3"
	GITHASH = "97e78a0"
	BUILT   = "2022-01-28T02:31:52+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
