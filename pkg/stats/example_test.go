package stats

import (
	"fmt"

	"github.com/Yessentemir/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1111,
			Category: "auto",
			Amount:   10_000,
			Status:   "OK",
		},
		{
			ID:       2222,
			Category: "auto",
			Amount:   20_000,
			Status:   "OK",
		},
		{
			ID:       3333,
			Category: "auto",
			Amount:   30_000,
			Status:   "OK",
		},
	}

	fmt.Println(Avg(payments))
	//Output: 20000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1111,
			Category: "auto",
			Amount:   10_000,
			Status:   "OK",
		},
		{
			ID:       2222,
			Category: "technology",
			Amount:   20_000,
			Status:   "OK",
		},
		{
			ID:       3333,
			Category: "auto",
			Amount:   30_000,
			Status:   "OK",
		},
	}

	var category types.Category = "auto"

	fmt.Println(TotalInCategory(payments, category))
	//Output: 40000

}
