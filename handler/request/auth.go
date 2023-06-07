package request

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FullName        string `json:"fullname"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
}
