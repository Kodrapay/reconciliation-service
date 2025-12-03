package dto

type ReconRunRequest struct {
	Source string `json:"source"`
	Date   string `json:"date"`
}

type ReconRunResponse struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Status string `json:"status"`
}
