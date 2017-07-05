package uds

type SID uint8

const (
	DiagnosticSessionControl        SID = 0x10
	ECUReset                        SID = 0x11
	ClearDiagnosticInformation      SID = 0x14
	ReadDTCInformation              SID = 0x19
	ReadDataByIdentifier            SID = 0x22
	ReadMemoryByAddress             SID = 0x23
	ReadScalingDataByIdentifier     SID = 0x24
	SecurityAccess                  SID = 0x27
	CommunicationControl            SID = 0x28
	ReadDataByPeriodicIdentifier    SID = 0x2A
	DynamicallyDefineDataIdentifier SID = 0x2C
	WriteDataByIdentifier           SID = 0x2E
	InputOutputControlByIdentifier  SID = 0x2F
	RoutineControl                  SID = 0x31
	RequestDownload                 SID = 0x34
	RequestUpload                   SID = 0x35
	TransferData                    SID = 0x36
	RequestTransferExit             SID = 0x37
	WriteMemoryByAddress            SID = 0x3D
	TesterPresent                   SID = 0x3E
	AccessTimingParameter           SID = 0x83
	SecuredDataTransmission         SID = 0x84
	ControlDTCSetting               SID = 0x85
	ResponseOnEvent                 SID = 0x86
	LinkControl                     SID = 0x87
)

type SessionType uint8

const (
	DefaultSession SessionType = 0x01
)
