package account_handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h AccountHandler) Upgrade(c *gin.Context) {
	cID, _ := c.Get(auth_entity.JwtClaim_UserID)

	userID := cID.((int64))

	err := h.accountUC.UpgradeAccount(userID)
	if err != nil {
		utils.GeneralErrorResponse(c, err, "Fail to upgrade account")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success upgrade account", nil)
}
