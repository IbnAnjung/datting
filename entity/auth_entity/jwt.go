package auth_entity

type UserJwtClaims struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	IsPremiumUser bool   `json:"is_premium_user"`
}

var (
	JwtClaim_UserID        = "UserID"
	JwtClaim_Username      = "Username"
	JwtClaim_IsPremiumUser = "IsPremiumUser"
)
