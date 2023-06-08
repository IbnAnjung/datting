package auth_entity

type RegisterInput struct {
	Username        string
	Password        string
	ConfirmPassword string
	FullName        string
	Age             int
	Gender          string
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
