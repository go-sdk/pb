package pb

var (
	VERSION = "dev"
	GITHASH = "4c2d804"
	BUILT   = "2022-02-09T18:04:34+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
