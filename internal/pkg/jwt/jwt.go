package jwt

import (
	"context"
	"encoding/json"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) uint {
	var uid uint
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = uint(int64Uid)
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

type JwtToken struct {
	AccessToken  string
	AccessExpire int64
	RefreshAfter int64
}

func GenerateToken(UserId uint, accessSecret string, accessExpire int64) (*JwtToken, error) {
	now := time.Now().Unix()
	accessToken, err := GetJwtToken(accessSecret, now, accessExpire, UserId)
	if err != nil {
		return nil, errors.Wrapf(code.ErrGenerateTokenError, "GetJwtToken err userId:%d , err:%v", UserId, err)
	}

	return &JwtToken{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func GetJwtToken(secretKey string, iat, seconds int64, userId uint) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["jwtUserId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}