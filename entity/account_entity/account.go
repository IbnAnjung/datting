package account_entity

type AccountUseCase interface {
	UpgradeAccount(userID int64) error
}
