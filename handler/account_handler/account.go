package account_handler

import (
	"github.com/IbnAnjung/datting/entity/account_entity"
)

type AccountHandler struct {
	accountUC account_entity.AccountUseCase
}

func NewAccountHandler(
	accountUC account_entity.AccountUseCase,
) AccountHandler {
	return AccountHandler{
		accountUC: accountUC,
	}
}
