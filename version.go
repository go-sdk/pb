package pb

var (
	VERSION = "dev"
	GITHASH = "ec94fd1"
	BUILT   = "2022-03-02T14:21:52+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
