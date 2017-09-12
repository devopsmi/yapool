package yapool

type Level uint

func (l Level) String() string {
	switch l {
	case Heartbeat:
		return HeartbeatInfo
	case Register:
		return RegisterInfo
	case Unregister:
		return UnregisterInfo
	case Warn:
		return WarnInfo
	case Error:
		return ErrorInfo
	case Fatal:
		return FatalInfo

	}
	return "[unknown]"
}

func IsLevel(l string) bool {
	for _, levelInfo := range AllLevelInfo {
		if levelInfo == l {
			return true
		}
	}
	return false
}

func SetLevel(l uint) Level {
	switch l {
	case 0:
		return Heartbeat
	case 1:
		return Register
	case 2:
		return Unregister
	case 3:
		return Warn
	case 4:
		return Error
	case 5:
		return Fatal
	}
	return Heartbeat
}

var AllLevelInfo = []string{
	HeartbeatInfo,
	RegisterInfo,
	UnregisterInfo,
	WarnInfo,
	ErrorInfo,
	FatalInfo,
}

const (
	Heartbeat Level = iota
	Register
	Unregister
	Warn
	Error
	Fatal
)

const (
	HeartbeatInfo  = "[heartbeat]"
	RegisterInfo   = "[register]"
	UnregisterInfo = "[unregister]"
	WarnInfo       = "[warn]"
	ErrorInfo      = "[error]"
	FatalInfo      = "[fatal]"
)
