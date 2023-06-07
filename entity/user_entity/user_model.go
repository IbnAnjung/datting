package user_entity

type UserModel struct {
	ID            int64
	Username      string
	Password      string
	Fullname      string
	Age           int
	Gender        string
	IsPremiumUser bool
}
