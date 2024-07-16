package saga

type SagaStatus string

const (
	STARTED      SagaStatus = "Started"
	FAILED       SagaStatus = "Failed"
	SUCCEEDED    SagaStatus = "Succeeded"
	PROCESSING   SagaStatus = "Processing"
	COMPENSATING SagaStatus = "Compensating"
	COMPENSATED  SagaStatus = "Compensated"
)
