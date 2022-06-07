package stats

import "github.com/ShavqatKavrakov/Lesson11_bank/v2/pkg/types"

//Avg рассчитывает среднюю сумму платёжа.
func Avg(payments []types.Payment) (sum types.Money) {
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) (sum types.Money) {
	for _, payment := range payments {
		if payment.Status == types.StatusInProgress {
			continue
		}
		if category != payment.Category {
			continue
		}
		sum += payment.Amount
	}
	return sum
}

//AvgInCategory находит среднюю сумму покупок в определённой категории.
func AvgInCategory(payments []types.Payment, category types.Category) (sum types.Money) {
	length := 0
	for _, payment := range payments {
		if category != payment.Category {
			continue
		}
		sum += payment.Amount
		length++
	}
	return sum / types.Money(length)
}

//CategoriesAvg считает среднюю сумму платёж по каждой категориии
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesAvg := map[types.Category]types.Money{}
	for _, payment := range payments {
		categoriesAvg[payment.Category] += payment.Amount
	}
	for key := range categoriesAvg {
		categoriesAvg[key] = AvgInCategory(payments, key)
	}
	return categoriesAvg
}
