package pb

var (
	VERSION = "v0.1.6"
	GITHASH = "80be8ed"
	BUILT   = "2022-01-28T16:05:04+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
