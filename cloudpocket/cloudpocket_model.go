package cloudpocket

type Model struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Budget      float64 `json:"budget"`
	Balance     float64 `json:"balance"`
	IsDefault   bool    `json:"isDefault"`
	Description string  `json:"description"`
	Current     string  `json:"current"`
	AccountId   int64   `json:"accountId"`
}
