package types

type OperationInputCLI struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int64   `json:"quantity"`
}

type OperationInput struct {
	OperationType string
	UnitCost      float64
	Quantity      int64
}
