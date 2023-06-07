package auth_entity

type RegisterInput struct {
	Username        string `json:"username" validate:"required,alphanumunicode,min=3,max=25"`
	Password        string `json:"password" validate:"required,min=5,max=50"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password,min=5,max=50"`
	FullName        string `json:"fullname" validate:"required,min=5,max=100"`
	Age             int    `json:"age" validate:"required,numeric,min=1,max=200"`
	Gender          string `json:"gender" validate:"required,gender"`
}

type RegisterOutput struct {
	ID            int64
	Username      string
	Password      string
	FullName      string
	Age           int
	Gender        string
	IsPremiumUser bool
}
