package controller

type RequestBalance struct {
	UserID uint    `json:"user_id"`
	AddSum float64 `json:"add_sum"`
}

type User struct {
	ID      uint    `json:"id"`
	Balance float64 `json:"balance"`
	Reserve float64 `json:"reserve"`
}

type Order struct {
	ID        uint    `json:"id"`
	ServiceID uint    `json:"service_id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
}
