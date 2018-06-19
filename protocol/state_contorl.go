package protocol

type StateControl []byte

var (
	OpenState     StateControl = []byte("O\r")
	LoopbackState StateControl = []byte("L\r")
	ResetState    StateControl = []byte("R\r")
)

func StateMsg(stat StateControl) []byte {
	return stat
}
