package version1

type SessionStatus string

const (
	SessionStatusCreated SessionStatus = "created"
	SessionStatusInUse   SessionStatus = "in_use"
	SessionStatusEnded   SessionStatus = "ended"
)

func SessionStatusFromString(status string) (SessionStatus, bool) {
	switch status {
	case string(SessionStatusCreated):
		return SessionStatusCreated, true
	case string(SessionStatusEnded):
		return SessionStatusEnded, true
	case string(SessionStatusInUse):
		return SessionStatusInUse, true
	default:
		return "", false
	}
}
