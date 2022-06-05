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
