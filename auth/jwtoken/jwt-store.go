package jwtoken

import (
	"fmt"

	"github.com/superryanguo/lightning/basic/cache"
)

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
