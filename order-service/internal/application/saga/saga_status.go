package saga

type SagaStatus string

const (
	STARTED      SagaStatus = "STARTED"
	FAILED       SagaStatus = "FAILED"
	SUCCEEDED    SagaStatus = "SUCCEEDED"
	PROCESSING   SagaStatus = "PROCESSING"
	COMPENSATING SagaStatus = "COMPENSATING"
	COMPENSATED  SagaStatus = "COMPENSATED"
)
