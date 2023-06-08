package user_entity

import "context"

const (
	DAILY_USER_VIEW_PROFILES string = "daily_user_view_profiles:"
)

type UserCacheRepository interface {
	SetDailyUserViewProfile(ctx context.Context, userID, viewedUserId int64) error
	GetDailyUserViewProfile(ctx context.Context, userID int64) ([]int64, error)
}
