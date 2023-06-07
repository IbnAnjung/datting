package auth_entity

type LoginInput struct {
	Username string `json:"username" validate:"required,alphanumunicode,min=3,max=25"`
	Password string `json:"password" validate:"required,min=5,max=50"`
}

type LoginOutput struct {
	ID            int64
	Username      string
	FullName      string
	Age           int
	Gender        string
	IsPremiumUser bool
	JwtToken      string
}
