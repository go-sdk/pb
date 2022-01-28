package pb

var (
	VERSION = "dev"
	GITHASH = "b0a406f"
	BUILT   = "2022-01-28T03:09:53+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
