package user_swap

import (
	"context"
	"log"

	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/user_swap_entity"
	"github.com/IbnAnjung/datting/utils"
)

func (uc UserSwapUC) SwapUserProfile(ctx context.Context, input user_swap_entity.SwapUserProfileInput) error {
	if err := uc.validator.ValidateStruct(input); err != nil {
		return err
	}

	authUser := user_entity.UserModel{}
	users, err := uc.userRepository.FindUserByIds([]int64{input.AuthUserID, input.SwappedProfileUserID})
	if err != nil {
		log.Printf("fail get data users %s", err.Error())
		return utils.ServerError{}
	}

	if len(users) != 2 {
		e := utils.DataNotFoundError
		e.Message = "user does not exists"
		return e
	}

	for _, v := range users {
		if v.ID == input.AuthUserID {
			authUser = v
			break
		}
	}

	swapUserIDs, err := uc.userCacheRepository.GetDailyUserSwapProfile(ctx, input.AuthUserID)
	if err != nil {
		return utils.ServerError{}
	}

	for _, id := range swapUserIDs {
		if id == input.SwappedProfileUserID {
			e := utils.UnprocessableEntityError
			e.Message = "you already react to this profile for today"
			return e
		}
	}

	if !authUser.IsPremiumUser && len(swapUserIDs) == 10 {
		e := utils.UnprocessableEntityError
		e.Message = "you have reached the maximum react limit, please upgrade to remove the limit"
		return e
	}

	if input.SwapType == user_swap_entity.SwapToLike {
		err = uc.userSwapRepository.SwapToLike(ctx, input.AuthUserID, input.SwappedProfileUserID)
	} else {
		err = uc.userSwapRepository.SwapToSkip(ctx, input.AuthUserID, input.SwappedProfileUserID)
	}

	if err != nil {
		return utils.ServerError{}
	}

	err = uc.userCacheRepository.SetDailyUserSwapProfile(ctx, input.AuthUserID, input.SwappedProfileUserID)
	if err != nil {
		return utils.ServerError{}
	}

	return nil
}
