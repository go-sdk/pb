package pb

var (
	VERSION = "dev"
	GITHASH = "0000a00"
	BUILT   = "2022-03-07T18:10:48+0800"
)

func VersionInfoMap() map[string]interface{} {
	return map[string]interface{}{
		"version": VERSION + "-" + GITHASH,
		"built":   BUILT,
	}
}
