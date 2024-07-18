package outbox

type OutboxStatus string

const (
	STARTED   OutboxStatus = "STARTED"
	COMPLETED OutboxStatus = "COMPLETED"
	FAILED    OutboxStatus = "FAILED"
)
