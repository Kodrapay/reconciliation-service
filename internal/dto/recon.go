package dto

type ReconRunRequest struct {
	Source string `json:"source"`
	Date   string `json:"date"` // YYYY-MM-DD in most flows; kept as string to stay flexible
}

type ReconRunResponse struct {
	ID        int    `json:"id"`
	Source    string `json:"source"`
	Date      string `json:"date"`
	Status    string `json:"status"`
	Matched   int    `json:"matched"`
	Unmatched int    `json:"unmatched"`
	CreatedAt string `json:"created_at"`
}
