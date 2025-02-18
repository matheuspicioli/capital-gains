package operation

import (
	"capital-gains/types"
	"math"
)

const (
	MINIMUN_TAX    = 20_000
	TAX_PERCENTAGE = 0.2
)

type OperationsManager struct {
	Stocks     int64
	Loss       float64
	TotalValue float64
}

type Transaction struct {
	UnitCost float64
	Quantity int64
}

type Tax struct {
	Tax float64
}

func FromOperationsInputCLI(operations []types.OperationInputCLI) []types.OperationInput {
	operationsInput := make([]types.OperationInput, len(operations))

	for i, operation := range operations {
		operationsInput[i] = types.OperationInput{
			OperationType: operation.Operation,
			UnitCost:      operation.UnitCost,
			Quantity:      operation.Quantity,
		}
	}

	return operationsInput
}

func Process(operations []types.OperationInput) []float64 {
	managementOperations := &OperationsManager{
		Stocks:     0,
		Loss:       0.0,
		TotalValue: 0.0,
	}

	taxes := make([]float64, len(operations))
	for i, operation := range operations {
		if operation.OperationType == "buy" {
			managementOperations.ExecuteBuy(Transaction{UnitCost: operation.UnitCost, Quantity: operation.Quantity})
			taxes[i] = 0.0
			continue
		}

		taxes[i] = managementOperations.ExecuteSell(Transaction{UnitCost: operation.UnitCost, Quantity: operation.Quantity})
	}

	return taxes
}

func (o *OperationsManager) ExecuteBuy(transaction Transaction) {
	o.Stocks += transaction.Quantity
	o.TotalValue += float64(transaction.Quantity) * transaction.UnitCost
}

func (o *OperationsManager) ExecuteSell(transaction Transaction) float64 {
	gain := o.Gain(transaction)

	if gain < 0 {
		o.Loss += math.Abs(gain)
		o.DeduceValues(transaction.Quantity, o.Average())
		return 0.0
	}

	if (transaction.UnitCost * float64(transaction.Quantity)) <= MINIMUN_TAX {
		if o.Loss > 0.0 {
			o.Loss -= gain
		}
		o.DeduceValues(transaction.Quantity, o.Average())
		return 0.0
	}

	o.DeduceValues(transaction.Quantity, o.Average())

	if o.Loss > 0.0 {
		if gain <= o.Loss {
			// think about: if gain >= MINIMUN_TAX, should it pay tax?
			o.Loss -= gain
			return 0.0
		}

		remainingGain := gain - o.Loss
		o.Loss = 0.0
		return o.CalculateTax(remainingGain)
	}

	return o.CalculateTax(gain)
}

func (o *OperationsManager) CalculateTax(gain float64) float64 {
	return gain * TAX_PERCENTAGE
}

func (o *OperationsManager) Average() float64 {
	return o.TotalValue / float64(o.Stocks)
}

func (o *OperationsManager) Gain(transaction Transaction) float64 {
	return o.UnitGain(transaction.UnitCost) * float64(transaction.Quantity)
}

func (o *OperationsManager) UnitGain(unitCost float64) float64 {
	return unitCost - o.Average()
}

func (o *OperationsManager) DeduceValues(quantity int64, average float64) {
	o.Stocks -= quantity
	o.TotalValue -= average * float64(quantity)
}
