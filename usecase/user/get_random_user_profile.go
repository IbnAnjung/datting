package user

import (
	"context"

	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/utils"
)

func (uc UserUC) GetRandomUserProfile(ctx context.Context, input user_entity.DetailUserInput) (user_entity.DetailUserOutput, error) {
	authUser, err := uc.userRepository.FindUserById(input.AuthUserID)
	if err != nil {
		return user_entity.DetailUserOutput{}, err
	}

	if authUser.ID == 0 {
		return user_entity.DetailUserOutput{}, utils.DataNotFoundError
	}

	viewedUserIDs, err := uc.userCacheRepository.GetDailyUserViewProfile(ctx, input.AuthUserID)
	if err != nil {
		return user_entity.DetailUserOutput{}, utils.ServerError{}
	}

	oppositeGender := utils.GetGenderOpposite(user_entity.Gender(authUser.Gender))
	user, err := uc.userRepository.FindUser(string(oppositeGender), viewedUserIDs)
	if err != nil {
		return user_entity.DetailUserOutput{}, utils.DataNotFoundError
	}

	if err := uc.userCacheRepository.SetDailyUserViewProfile(ctx, input.AuthUserID, user.ID); err != nil {
		return user_entity.DetailUserOutput{}, utils.ServerError{}
	}

	return user_entity.DetailUserOutput{
		ID:            user.ID,
		Username:      user.Username,
		Fullname:      user.Fullname,
		Age:           user.Age,
		Gender:        user.Gender,
		IsPremiumUser: user.IsPremiumUser,
	}, nil
}
