package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hyahm/goconfig"
)

// 2. 自定义Claims（包含标准Claims + 自定义字段）
// jwt.RegisteredClaims 是v5版本的标准Claims，包含过期时间、签发时间等
type CustomClaims struct {
	Uid                  int64 `json:"uid"` // 自定义字段：用户ID
	jwt.RegisteredClaims       // 嵌入标准Claims，必须
}

// 3. 生成JWT Token
func GenerateToken(userID int64) (string, error) {
	// 设置Token有效期（例如2小时）
	expireTime := time.Now().Add(goconfig.ReadDuration("expiration"))

	// 构造自定义Claims
	claims := CustomClaims{
		Uid: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			Issuer:    "itflow",                       // 签发者（自定义）
			Subject:   "Authorization",                // 主题（自定义）
		},
	}

	// 创建Token：指定签名算法和Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用密钥签名并生成最终的Token字符串
	tokenString, err := token.SignedString([]byte(goconfig.ReadString("salt")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 4. 验证并解析JWT Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析Token（第二个参数是回调函数，用于获取签名密钥）
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{}, // 传入自定义Claims的指针，用于接收解析结果
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法是否正确（防止算法伪造）
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回签名密钥
			return []byte(goconfig.ReadString("salt")), nil
		},
	)

	// 解析失败的通用错误
	if err != nil {
		return nil, err
	}

	// 验证Token有效且解析Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
