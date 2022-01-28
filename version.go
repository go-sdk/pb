package pb

var (
	VERSION = "dev"
	GITHASH = "b698a23"
	BUILT   = "2022-01-28T07:49:26+0000"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
