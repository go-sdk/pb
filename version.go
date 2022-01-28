package pb

var (
	VERSION = "dev"
	GITHASH = "4011613"
	BUILT   = "2022-01-28T16:03:24+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
