package pb

var (
	VERSION = "dev"
	GITHASH = "082455e"
	BUILT   = "2022-01-28T03:24:34+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
