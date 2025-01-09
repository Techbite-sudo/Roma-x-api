package game

// Payline represents a single payline pattern
type Payline struct {
	Number  int
	Pattern [][2]int // Array of [row, col] positions for each payline
}

// Initialize paylines patterns according to the game grid
var Paylines = []Payline{
	// Payline 1: Top horizontal
	{Number: 1, Pattern: [][2]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
	}},

	// Payline 2: Middle horizontal
	{Number: 2, Pattern: [][2]int{
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
	}},

	// Payline 3: Bottom horizontal
	{Number: 3, Pattern: [][2]int{
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
	}},

	// Payline 4: V shape top to middle
	{Number: 4, Pattern: [][2]int{
		{0, 0}, {1, 1}, {1, 2}, {1, 3}, {0, 4},
	}},

	// Payline 5: V shape middle to bottom
	{Number: 5, Pattern: [][2]int{
		{1, 0}, {2, 1}, {2, 2}, {2, 3}, {1, 4},
	}},

	// Payline 6: Zigzag top-middle
	{Number: 6, Pattern: [][2]int{
		{1, 0}, {0, 1}, {1, 2}, {0, 3}, {1, 4},
	}},

	// Payline 7: Zigzag middle-bottom
	{Number: 7, Pattern: [][2]int{
		{1, 0}, {2, 1}, {1, 2}, {2, 3}, {1, 4},
	}},

	// Payline 8: Zigzag top-middle downward
	{Number: 8, Pattern: [][2]int{
		{0, 0}, {1, 1}, {0, 2}, {1, 3}, {0, 4},
	}},

	// Payline 9: Zigzag middle-bottom downward
	{Number: 9, Pattern: [][2]int{
		{1, 0}, {2, 1}, {1, 2}, {2, 3}, {1, 4},
	}},

	// Payline 10: V shape middle
	{Number: 10, Pattern: [][2]int{
		{1, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 4},
	}},

	// Payline 11: Inverse V shape top
	{Number: 11, Pattern: [][2]int{
		{0, 0}, {1, 1}, {1, 2}, {1, 3}, {0, 4},
	}},

	// Payline 12: Inverse V shape bottom
	{Number: 12, Pattern: [][2]int{
		{1, 0}, {2, 1}, {2, 2}, {2, 3}, {1, 4},
	}},

	// Payline 13: V shape bottom
	{Number: 13, Pattern: [][2]int{
		{2, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 4},
	}},

	// Payline 14: Diamond shape
	{Number: 14, Pattern: [][2]int{
		{0, 0}, {1, 1}, {2, 2}, {1, 3}, {0, 4},
	}},

	// Payline 15: Inverse diamond shape
	{Number: 15, Pattern: [][2]int{
		{2, 0}, {1, 1}, {0, 2}, {1, 3}, {2, 4},
	}},
}

// Helper function to check symbols along a payline
func CheckPayline(grid Grid, payline Payline) (string, int) {
	symbols := make([]string, 5)

	// Get symbols along payline
	for i, pos := range payline.Pattern {
		symbols[i] = grid.Symbols[pos[1]][pos[0]]
	}

	// Count matching symbols from left
	firstSymbol := symbols[0]
	count := 1

	for i := 1; i < len(symbols); i++ {
		if symbols[i] == firstSymbol || symbols[i] == "Wild" || (firstSymbol == "Wild" && i == 1) {
			count++
		} else {
			break
		}
	}

	if count < 3 {
		return "", 0
	}

	if firstSymbol == "Wild" {
		// Find the first non-Wild symbol to determine the winning symbol type
		for _, sym := range symbols[1:count] {
			if sym != "Wild" {
				return sym, count
			}
		}
		// If all are Wilds, count as highest paying symbol
		return "Helmet", count // Helmet is highest paying symbol
	}

	return firstSymbol, count
}
