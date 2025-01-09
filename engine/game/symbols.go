package game

func InitializeSymbols() map[string]Symbol {
	return map[string]Symbol{
		"Helmet": {
			Name: "Helmet",
			Payouts: map[int]int{
				3: 18,
				4: 180,
				5: 1200,
			},
		},
		"Eye": {
			Name: "Eye",
			Payouts: map[int]int{
				3: 18,
				4: 120,
				5: 600,
			},
		},
		"Lion": {
			Name: "Lion",
			Payouts: map[int]int{
				3: 12,
				4: 60,
				5: 240,
			},
		},
		"Sword": {
			Name: "Sword",
			Payouts: map[int]int{
				3: 12,
				4: 36,
				5: 144,
			},
		},
		"Shield": {
			Name: "Shield",
			Payouts: map[int]int{
				3: 6,
				4: 18,
				5: 90,
			},
		},
		"Cup": {
			Name: "Cup",
			Payouts: map[int]int{
				3: 6,
				4: 12,
				5: 36,
			},
		},
		"Grapes": {
			Name: "Grapes",
			Payouts: map[int]int{
				3: 6,
				4: 12,
				5: 36,
			},
		},
	}
}
