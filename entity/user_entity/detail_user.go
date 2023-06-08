package user_entity

type DetailUserInput struct {
	AuthUserID int64
}

type DetailUserOutput struct {
	ID            int64
	Username      string
	Fullname      string
	Age           int
	Gender        string
	IsPremiumUser bool
}
