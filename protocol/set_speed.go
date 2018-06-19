package protocol

type BusSpeed []byte

var (
	S10K  BusSpeed = []byte("S0\r")
	S20K  BusSpeed = []byte("S1\r")
	S50K  BusSpeed = []byte("S2\r")
	S100K BusSpeed = []byte("S3\r")
	S125K BusSpeed = []byte("S4\r")
	S250K BusSpeed = []byte("S5\r")
	S500K BusSpeed = []byte("S6\r")
	S800K BusSpeed = []byte("S7\r")
	S1M   BusSpeed = []byte("S8\r")
)

func SetSpeedMsg(baudRate BusSpeed) []byte {
	return baudRate
}
