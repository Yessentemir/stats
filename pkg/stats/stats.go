package stats

import (
	"github.com/Yessentemir/bank/v2/pkg/types"
)

// Avg рассчитывет среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	quantity := len(payments)
	var average types.Money
	for _, payments := range payments {

		if payments.Amount < 0 {
			continue
		}

		if payments.Status == "FAIL" {
			continue
		}

		sum += payments.Amount
		average = sum / types.Money(quantity)

	}
	return average
}

// TotalinCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payments := range payments {
		if payments.Amount < 0 {
			continue
		}

		if category != payments.Category {
			continue
		}

		if payments.Status == "FAIL" {
			continue
		}
		sum += payments.Amount

	}
	return sum
}
