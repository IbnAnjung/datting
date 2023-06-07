package response

type RegisterResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	JwtToken string `json:"jwt_token"`
}
