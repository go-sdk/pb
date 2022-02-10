package pb

var (
	VERSION = "v0.2.1"
	GITHASH = "65a6809"
	BUILT   = "2022-02-10T09:51:04+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
