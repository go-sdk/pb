package pb

var (
	VERSION = "latest"
	GITHASH = "c371f03"
	BUILT   = "2022-01-28T02:11:39+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
