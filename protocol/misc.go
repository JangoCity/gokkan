package protocol

type Misc []byte

var (
	VersionHigh Misc = []byte("V\r")
	Serial      Misc = []byte("N\r")
	TSOn        Misc = []byte("Z0\r")
	TSOff       Misc = []byte("Z1\r")
	ReadStatus       Misc = []byte("F\r")
)

func MiscMsg(stat Misc) []byte {
	return stat
}
