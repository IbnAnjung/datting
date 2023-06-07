package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/IbnAnjung/datting/driver"
	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf, _ := LoadConfig()

		jwt := driver.NewJwt(conf.App.Name, conf.Jwt.SecretKey, conf.Jwt.ExpireDuration)

		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "missing authorization")
			c.Abort()
			return
		}

		authPart := strings.Split(auth, " ")
		bearer := strings.ToLower(authPart[0])

		if !strings.EqualFold(bearer, "bearer") || len(authPart) != 2 {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid authorization format")
			c.Abort()
			return
		}

		token := authPart[1]
		claim, err := jwt.ParseToken(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		claimMap, ok := claim.(map[string]interface{})
		if !ok {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}
		var userJwtClaims auth_entity.UserJwtClaims
		jsonClaims, _ := json.Marshal(claimMap)

		json.Unmarshal(jsonClaims, &userJwtClaims)

		c.Set(auth_entity.JwtClaim_UserID, userJwtClaims.ID)
		c.Set(auth_entity.JwtClaim_Username, userJwtClaims.Username)
		c.Set(auth_entity.JwtClaim_IsPremiumUser, userJwtClaims.IsPremiumUser)

		c.Next()

	}
}
