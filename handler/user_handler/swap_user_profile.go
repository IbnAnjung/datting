package user_handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_swap_entity"
	"github.com/IbnAnjung/datting/handler/user_handler/dto"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) SwapUserProfile(c *gin.Context) {
	var request dto.SwapUserProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorValidationResponse(c, err.Error(), []string{})
		return
	}

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

	err := h.userSwapUC.SwapUserProfile(c.Request.Context(), user_swap_entity.SwapUserProfileInput{
		AuthUserID:           userId,
		SwappedProfileUserID: request.UserID,
		SwapType:             user_swap_entity.SwapType(request.Type),
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err, "Fail to react profile")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "success swap", nil)
}
