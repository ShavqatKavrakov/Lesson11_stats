package stats

import (
	"fmt"

	"github.com/ShavqatKavrakov/Lesson11_bank/pkg/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
		{
			Amount: 12_000_00,
		},
	}))

	//Output:
	//500000
}
func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   1_000_00,
			Category: "Restaurant",
		},
		{
			Amount:   2_000_00,
			Category: "Car",
		},
		{
			Amount:   12_000_00,
			Category: "Restaurant",
		},
		{
			Amount:   14_000_00,
			Category: "Car",
		},
	}, "Restaurant"))

	//Output:
	//1300000
}
