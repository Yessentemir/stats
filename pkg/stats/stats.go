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

// TotalinCategory2 находит сумму покупок в определенной категории.
func TotalInCategory2(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {

		if payment.Category != category {
			continue
		}

		sum += payment.Amount

	}
	return sum
}

// QuantityOfCategoriesInPayment считает количество категорий в платеже.
func QuantityOfCategoriesInPayment(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sum++
		}

	}
	return sum
}

// CategoriesAvg cчитает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] = TotalInCategory2(payments, payment.Category) / QuantityOfCategoriesInPayment(payments, payment.Category)
	}

	return categories
}

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}

	for i, value := range second {
		result[i] = value

		if el, ok := first[i]; ok {
			result[i] -= el
		}
	}

	for i, value := range first {
		if _, ok := result[i]; !ok {
			result[i] = value * (-1)
		}
	}

	return result
}
