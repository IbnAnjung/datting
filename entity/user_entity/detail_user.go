package user_entity

type DetailUserInput struct {
	AuthUserID int64 `json:"id"`
}

type DetailUserOutput struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	IsPremiumUser bool   `json:"is_premium"`
}
