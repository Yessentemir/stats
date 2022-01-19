package stats

import (
	"fmt"

	"github.com/Yessentemir/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1111,
			Category: "auto",
			Amount:   10_000,
		},
		{
			ID:       2222,
			Category: "auto",
			Amount:   20_000,
		},
		{
			ID:       3333,
			Category: "auto",
			Amount:   30_000,
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
		},
		{
			ID:       2222,
			Category: "technology",
			Amount:   20_000,
		},
		{
			ID:       3333,
			Category: "auto",
			Amount:   30_000,
		},
	}

	var category types.Category = "auto"

	fmt.Println(TotalInCategory(payments, category))
	//Output: 40000

}
