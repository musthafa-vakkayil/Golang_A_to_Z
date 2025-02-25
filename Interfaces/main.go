package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

// ?

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}

	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	// ?

	switch ty := n.(type) {
	case directMessage:
		return ty.senderUsername, ty.importance()
	case groupMessage:
		return ty.groupName, ty.importance()
	case systemAlert:
		return ty.alertCode, ty.importance()
	default:
		return "", 0
	}
}
