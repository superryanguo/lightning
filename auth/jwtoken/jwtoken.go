package jwtoken

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/config"
)

var (
	//30 days
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "superryan.guo.lightning.topic.auth.tokenExpired"
)

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func MakeAccessToken(subject *Subject) (ret string, err error) {
	m, err := createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] fail，err: %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] fail，err: %s", err)
	}

	err = saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] save token failure，err: %s", err)
	}

	return
}

func GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] fail，err: %s", err)
	}

	return
}

func DelUserAccessToken(tk string) (err error) {
	claims, err := ParseToken(tk)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] wrong token，err: %s", err)
	}

	err = delTokenFromCache(&Subject{
		ID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] clear token，err: %s", err)
	}

	//broadcast
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}
	if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
		log.Errorf("[pub token fail] %v", err)
	} else {
		fmt.Println("[pub token del]", string(msg.Body))
	}

	return
}
func createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "superryan.guo.lightning",
		Subject:   subject.ID,
	}

	return
}
func ParseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("illegal token format: %v", token.Header["alg"])
		}
		return []byte(config.GetJwtConfig().GetSecretKey()), nil
	})

	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[ParseToken] expired token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[ParseToken] illegal token, err:%s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[ParseToken] illegal token")
	}

	return mapClaimToJwClaim(claims), nil
}

func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}

	return jC
}
