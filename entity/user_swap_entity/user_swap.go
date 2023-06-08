package user_swap_entity

import "context"

type UserSwapModel struct {
	UserID        int64
	SwappedUserId int64
	PassCount     int64
	LikeCount     int64
}

type SwapType string

const (
	SwapToLike SwapType = "like"
	SwapToPass SwapType = "pass"
)

type UserSwapUseCase interface {
	SwapUserProfile(ctx context.Context, input SwapUserProfileInput) error
}

type UserSwapRepository interface {
	SwapToLike(ctx context.Context, userId, swappedUserId int64) error
	SwapToSkip(ctx context.Context, userId, swappedUserId int64) error
}
