package enum

type OutboxStatus string

const (
	OUTBOX_STARTED   OutboxStatus = "STARTED"
	OUTBOX_COMPLETED OutboxStatus = "COMPLETED"
	OUTBOX_FAILED    OutboxStatus = "FAILED"
)
