package uds

type DiagnosticService struct {
}

func RequestForSession(sessionType SessionType) *SingleFrame {
	return &SingleFrame{DiagnosticSessionControl, []byte{byte(DefaultSession), 0x0, 0x0, 0x0, 0x0, 0x0}, 2}
}
