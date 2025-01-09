package game

type Symbol struct {
	Name    string
	Payouts map[int]int
}

type Grid struct {
	Symbols [][]string
}

type SpinRequest struct {
	Bet      float64 `json:"bet"`
	PlayerID string  `json:"player_id"`
}

type SpinResponse struct {
	Grid         Grid      `json:"grid"`
	WinningLines []WinLine `json:"winning_lines"`
	TotalWin     float64   `json:"total_win"`
	FreeGames    int       `json:"free_games"`
	IsBonusGame  bool      `json:"is_bonus_game"`
}

type WinLine struct {
	PaylineNumber int     `json:"payline_number"`
	Symbols       string  `json:"symbols"`
	Count         int     `json:"count"`
	Payout        float64 `json:"payout"`
}
