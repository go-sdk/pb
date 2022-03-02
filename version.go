package pb

var (
	VERSION = "dev"
	GITHASH = "14dd854"
	BUILT   = "2022-03-02T14:00:23+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
