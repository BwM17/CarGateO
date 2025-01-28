package dtos

type CarDto struct {
	Model       string `json:"model"`
	Numberplate string `json:"numberplate"`
	Owner       string `json:"owner"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	LeasedAt    string `json:"lease"`
	ReturnAt    string `json:"return"`
}