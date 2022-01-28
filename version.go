package pb

var (
	VERSION = "latest"
	GITHASH = "697bb32"
	BUILT   = "2022-01-28T02:10:53+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
