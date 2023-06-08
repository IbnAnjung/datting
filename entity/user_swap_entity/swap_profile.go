package user_swap_entity

type SwapUserProfileInput struct {
	AuthUserID           int64    `json:"auth_user_id" validate:"required,numeric,min=1"`
	SwappedProfileUserID int64    `json:"swapped_profile_user_id" validate:"required,numeric,min=1,nefield=AuthUserID"`
	SwapType             SwapType `json:"swap_type" validate:"required,swap_type"`
}
