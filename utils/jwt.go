package util

import (
	"errors"
	"strconv"

	"github.com/AstraProtocol/reward-libs/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromClaims(ctx *gin.Context) (uint, error) {
	jwtClaims, err := middleware.GetJWTClaims(ctx)
	if err != nil {
		return 0, errors.New("bad jwt claims")
	}
	uid, err := strconv.ParseUint(jwtClaims["sub"].(string), 10, 32)
	if err != nil {
		return 0, errors.New("bad sub claims")
	}
	return uint(uid), nil
}
