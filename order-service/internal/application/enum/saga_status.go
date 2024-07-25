package enum

type SagaStatus string

const (
	SAGA_STARTED      SagaStatus = "STARTED"
	SAGA_FAILED       SagaStatus = "FAILED"
	SAGA_SUCCEEDED    SagaStatus = "SUCCEEDED"
	SAGA_PROCESSING   SagaStatus = "PROCESSING"
	SAGA_COMPENSATING SagaStatus = "COMPENSATING"
	SAGA_COMPENSATED  SagaStatus = "COMPENSATED"
)
