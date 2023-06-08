package utils

import "github.com/IbnAnjung/datting/entity/user_entity"

func GetGenderOpposite(gender user_entity.Gender) user_entity.Gender {
	if gender == user_entity.MALE {
		return user_entity.FEMALE
	}

	return user_entity.MALE
}
