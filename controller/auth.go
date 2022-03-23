package controller

import (
	"NetworkDisk/data/config"
	"NetworkDisk/library"
	"NetworkDisk/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	UserId int `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 权限认证 - 拦截器
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		data := gin.H{
			"location": "/user/login",
		}

		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusOK, library.RespMessage(20001,"无法认证，重新登录", data))
			return
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		claims, err := ParseToken(auth)
		if err != nil {
			c.Abort()
			message := "认证权限已过期，请重新登录"
			c.JSON(http.StatusOK, library.RespMessage(20002,message, data))
			return
		} else {
			log.Println("token 正确")
		}

		// TODO token续签

		c.Set("claims", claims)
		c.Next()
	}
}

// 获取 token
func CreateToken(user model.User) string {
	// 过期时间
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	claims := MyClaims{
		user.Id,
		user.UserName,
		jwt.StandardClaims {
			ExpiresAt: expiresTime,       // 失效时间
			IssuedAt:  time.Now().Unix(), // 签发时间
			NotBefore: time.Now().Unix(), // 生效时间
			Issuer:    "gin hello",       // 签发人
		},
	}

	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		return "Bearer " + token
	} else {
		return ""
	}
}

// 解析 token
func ParseToken(token string) (*MyClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*MyClaims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 获取 token 内容
func GetToken(c *gin.Context) (*MyClaims, bool) {
	claims, bool := c.Get("claims")
	if !bool {
		log.Println("claims 不存在")
		return  nil, false
	}
	return claims.(*MyClaims), bool
}




