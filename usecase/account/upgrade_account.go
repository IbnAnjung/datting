package account

import "github.com/IbnAnjung/datting/utils"

func (uc AccountUC) UpgradeAccount(userID int64) error {
	user, err := uc.userRepository.FindUserById(userID)
	if err != nil {
		return err
	}

	if user.IsPremiumUser {
		e := utils.UnprocessableEntityError
		e.Message = "this account is already premium"
		return e
	}

	user.IsPremiumUser = true
	err = uc.userRepository.UpdateUser(&user)
	if err != nil {
		return err
	}
	return nil
}
