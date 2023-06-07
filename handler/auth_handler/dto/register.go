package dto

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FullName        string `json:"fullname"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
}

type RegisterResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	JwtToken string `json:"jwt_token"`
}
