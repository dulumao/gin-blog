package utils

import (
	"gin-blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

// JWT 参考：http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
// 自定义 Payload（负载）
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		jwt.StandardClaims{
			Issuer:    "gin-blog",        // 签发人
			ExpiresAt: expireTime.Unix(), // 过期时间
			Subject:   "api auth",        // 主题
			Audience:  "all",             // 受众
			NotBefore: nowTime.Unix(),    // 生效时间
			IssuedAt:  nowTime.Unix(),    //签发时间
			//Id:        nil,               // 编号
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析 token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 返回key
func keyFunc(token *jwt.Token) (interface{}, error) {
	return jwtSecret, nil
}
