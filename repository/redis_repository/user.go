package redis_repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/util_entity"
	"github.com/IbnAnjung/datting/utils"
)

type UserCacheRepository struct {
	cache util_entity.Caching
}

func NewUserCacheRepository(
	cache util_entity.Caching,
) UserCacheRepository {
	return UserCacheRepository{
		cache: cache,
	}
}

func (r UserCacheRepository) SetDailyUserViewProfile(ctx context.Context, userID, viewedUserId int64) error {
	now := time.Now()

	if err := r.cache.PushList(fmt.Sprintf("%s:%d", user_entity.DAILY_USER_VIEW_PROFILES, userID), viewedUserId).
		ExpireAt(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, time.Local)).Do(ctx); err != nil {
		log.Printf("fail set user daily view profile error: %s", err.Error())
		return utils.ServerError{}
	}

	return nil
}

func (r UserCacheRepository) GetDailyUserViewProfile(ctx context.Context, userID int64) ([]int64, error) {
	strIds, err := r.cache.GetList(ctx, fmt.Sprintf("%s:%d", user_entity.DAILY_USER_VIEW_PROFILES, userID), 0, -1)
	if err != nil {
		log.Printf("fail get user daily view profile error: %s", err.Error())
		return []int64{0}, utils.ServerError{}
	}

	ids := []int64{}
	for _, strId := range strIds {
		if id, err := strconv.ParseInt(strId, 10, 64); err == nil {
			ids = append(ids, id)
		} else {
			log.Printf("invalid user_id: %s", strId)
		}
	}

	return ids, nil
}

func (r UserCacheRepository) SetDailyUserSwapProfile(ctx context.Context, userID, swappedUserId int64) error {
	now := time.Now()

	if err := r.cache.PushList(fmt.Sprintf("%s:%d", user_entity.DAILY_USER_SWAP_PROFILES, userID), swappedUserId).
		ExpireAt(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, time.Local)).Do(ctx); err != nil {
		log.Printf("fail set user daily swap profile error: %s", err.Error())
		return utils.ServerError{}
	}

	return nil
}

func (r UserCacheRepository) GetDailyUserSwapProfile(ctx context.Context, userID int64) ([]int64, error) {
	strIds, err := r.cache.GetList(ctx, fmt.Sprintf("%s:%d", user_entity.DAILY_USER_SWAP_PROFILES, userID), 0, -1)
	if err != nil {
		log.Printf("fail get user daily swap profile error: %s", err.Error())
		return []int64{0}, utils.ServerError{}
	}

	ids := []int64{}
	for _, strId := range strIds {
		if id, err := strconv.ParseInt(strId, 10, 64); err == nil {
			ids = append(ids, id)
		} else {
			log.Printf("invalid user_id: %s", strId)
		}
	}

	return ids, nil
}
