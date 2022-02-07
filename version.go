package pb

var (
	VERSION = "dev"
	GITHASH = "b7526f9"
	BUILT   = "2022-02-07T14:08:58+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
