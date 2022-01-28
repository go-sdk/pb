package pb

var (
	VERSION = "v.."
	GITHASH = "b30a8fe"
	BUILT   = "2022-01-28T02:30:52+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
