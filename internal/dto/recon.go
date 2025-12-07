package dto

type ReconRunRequest struct {
	Source string `json:"source"`
	Date   string `json:"date"`
}

type ReconRunResponse struct {
	ID     int    `json:"id"`
	Source string `json:"source"`
	Status string `json:"status"`
}
