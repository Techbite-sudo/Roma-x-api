package service

import (
	"math/rand"
	"roma-x-api/engine/game"
	"time"
)

type GameService struct {
	symbols map[string]game.Symbol
}

func NewGameService(symbols map[string]game.Symbol) *GameService {
	rand.Seed(time.Now().UnixNano())
	return &GameService{
		symbols: symbols,
	}
}

func (s *GameService) GenerateGameOutcome(bet float64) game.SpinResponse {
	grid := s.generateGrid()
	isWin := rand.Float64() < 0.3

	var winningLines []game.WinLine
	var totalWin float64
	var freeGames int
	var isBonusGame bool

	if isWin {
		winningLines = s.generateWinningLines(grid, bet)

		for _, line := range winningLines {
			totalWin += line.Payout
		}

		if rand.Float64() < 0.03 {
			isBonusGame = true
		}

		if rand.Float64() < 0.05 {
			freeGames = s.generateFreeGames()
		}

		if totalWin > game.MaxPayout {
			totalWin = game.MaxPayout
		}
	}

	return game.SpinResponse{
		Grid:         grid,
		WinningLines: winningLines,
		TotalWin:     totalWin,
		FreeGames:    freeGames,
		IsBonusGame:  isBonusGame,
	}
}

func (s *GameService) generateGrid() game.Grid {
	symbolList := []string{"Helmet", "Eye", "Lion", "Sword", "Shield", "Cup", "Grapes", "Wild", "Bonus"}
	grid := make([][]string, 5)

	for i := range grid {
		grid[i] = make([]string, 3)
		for j := range grid[i] {
			grid[i][j] = symbolList[rand.Intn(len(symbolList))]
		}
	}

	return game.Grid{Symbols: grid}
}

func (s *GameService) generateWinningLines(grid game.Grid, bet float64) []game.WinLine {
	var winningLines []game.WinLine

	// Randomly select number of winning lines (1-3)
	numWins := rand.Intn(3) + 1

	// Keep track of used paylines to avoid duplicates
	usedPaylines := make(map[int]bool)

	for i := 0; i < numWins; i++ {
		// Generate a random winning line
		winLine := s.generateSingleWinningLine(grid, bet, usedPaylines)
		if winLine != nil {
			winningLines = append(winningLines, *winLine)
			usedPaylines[winLine.PaylineNumber] = true
		}
	}

	return winningLines
}

func (s *GameService) generateSingleWinningLine(grid game.Grid, bet float64, usedPaylines map[int]bool) *game.WinLine {
	// Try up to 5 times to find an unused payline
	for attempt := 0; attempt < 5; attempt++ {
		// Select random payline number (1-15)
		paylineNum := rand.Intn(15) + 1

		if usedPaylines[paylineNum] {
			continue
		}

		// Get symbols along this payline and use them
		symbols := s.getPaylineSymbols(grid, game.Paylines[paylineNum-1])

		// Get a random symbol from the first 3 positions to ensure left-to-right wins
		winningSymbol := symbols[rand.Intn(3)]

		// Skip if the winning symbol is "Wild" or "Bonus"
		if winningSymbol == "Wild" || winningSymbol == "Bonus" {
			continue
		}

		// Generate winning combination
		symbolCount := rand.Intn(3) + 3 // 3-5 matching symbols

		// Calculate payout
		if payout, ok := s.symbols[winningSymbol].Payouts[symbolCount]; ok {
			return &game.WinLine{
				PaylineNumber: paylineNum,
				Symbols:       winningSymbol,
				Count:         symbolCount,
				Payout:        float64(payout) * bet / game.PaylineCount,
			}
		}
	}

	return nil
}

func (s *GameService) getPaylineSymbols(grid game.Grid, payline game.Payline) []string {
	symbols := make([]string, len(payline.Pattern))
	for i, pos := range payline.Pattern {
		symbols[i] = grid.Symbols[pos[1]][pos[0]]
	}
	return symbols
}

func (s *GameService) generateFreeGames() int {
	// Possible free game amounts based on number of matching symbols
	freeGames := []int{3, 5, 10, 20}
	return freeGames[rand.Intn(len(freeGames))]
}

// func (s *GameService) checkForBonusSymbols(grid game.Grid) bool {
// 	bonusCount := 0
// 	for i := range grid.Symbols {
// 		for j := range grid.Symbols[i] {
// 			if grid.Symbols[i][j] == "Bonus" {
// 				bonusCount++
// 				if bonusCount >= 3 {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// func (s *GameService) calculateBonusWin(bet float64, choice string) float64 {
// 	multipliers := map[string]float64{
// 		"dual_swords":  20,
// 		"single_sword": 10,
// 		"lion":         3,
// 	}

// 	if multiplier, ok := multipliers[choice]; ok {
// 		return bet * multiplier
// 	}
// 	return 0
// }
