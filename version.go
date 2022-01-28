package pb

var (
	VERSION = "dev"
	GITHASH = "4d00357"
	BUILT   = "2022-01-28T16:10:45+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
