package jwtoken

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/basic/config"
)

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

func saveTokenToCache(subject *Subject, val string) (err error) {
	if err = cache.SaveToCacheDate(tokenIDKeyPrefix+subject.ID, []byte(val), tokenExpiredDate); err != nil {
		return fmt.Errorf("[saveTokenToCache] fail err:" + err.Error())
	}
	return
}

func delTokenFromCache(subject *Subject) (err error) {
	if err = cache.DelFromCache(tokenIDKeyPrefix + subject.ID); err != nil {
		return fmt.Errorf("[delTokenFromCache] fail err:" + err.Error())
	}
	return
}

func getTokenFromCache(subject *Subject) (token string, err error) {
	tokenCached, err := cache.GetFromCache(tokenIDKeyPrefix + subject.ID)
	if err != nil {
		return token, fmt.Errorf("[getTokenFromCache] token not exist %s", err)
	}

	return string(tokenCached), nil
}

func parseToken(tk string) (c *jwt.StandardClaims, err error) {
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
				return nil, fmt.Errorf("[parseToken] expired token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] illegal token, err:%s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] illegal token")
	}

	return mapClaimToJwClaim(claims), nil
}

func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{
		Subject: claims["sub"].(string),
	}

	return jC
}
