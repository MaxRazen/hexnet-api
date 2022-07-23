package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"hexnet/api/common"
	"time"
)

var middleware *jwt.GinJWTMiddleware

const (
	identityKey = "sub"
	audienceKey = "aud"
)

func NewAuthMiddleware() *jwt.GinJWTMiddleware {
	conf := common.GetConfig()

	m, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       conf.AppName,
		IdentityKey: identityKey,
		Key:         []byte(conf.Env.Auth.Secret),
		Timeout:     time.Duration(int64(time.Minute) * int64(conf.Env.Auth.Timeout)),
		MaxRefresh:  time.Duration(int64(time.Minute) * int64(conf.Env.Auth.Refresh)),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			d, _ := data.(jwt.MapClaims)
			return d
		},
	})

	if err != nil {
		panic("AuthMiddleware could not be initialized")
	}

	middleware = m

	return middleware
}
