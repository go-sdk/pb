package pb

var (
	VERSION = "latest"
	GITHASH = "58f0b5e"
	BUILT   = "2022-01-28T02:17:32+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
