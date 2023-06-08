package user_swap_entity

type SwapUserProfileInput struct {
	AuthUserID           int64
	SwappedProfileUserID int64
	SwapType             SwapType
}
