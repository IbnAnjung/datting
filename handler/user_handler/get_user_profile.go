package user_handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/handler/user_handler/dto"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) GetUserProfile(c *gin.Context) {
	id, ok := c.Get(auth_entity.JwtClaim_UserID)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	userId, ok := id.(int64)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	user, err := h.userUC.GetRandomUserProfile(c.Request.Context(), user_entity.DetailUserInput{
		AuthUserID: userId,
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err, "Fail to show profile")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "get profile success", dto.DetailUserResponse{
		ID:            user.ID,
		FullName:      user.Fullname,
		Age:           user.Age,
		Gender:        user.Gender,
		IsPremiumUser: user.IsPremiumUser,
	})
}
