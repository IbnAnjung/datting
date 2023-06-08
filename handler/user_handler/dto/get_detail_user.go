package dto

type DetailUserResponse struct {
	ID            int64  `json:"id"`
	FullName      string `json:"fullname"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	IsPremiumUser bool   `json:"is_premium"`
}
