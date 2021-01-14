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
	claims, err := parseToken(tk)
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
