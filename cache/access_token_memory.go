package cache

import (
	"errors"
	"github.com/Unknwon/goconfig"
	"github.com/sirupsen/logrus"
	"time"
)

type accessTokenCache struct {
	AppID   uint      //应用id
	Expires time.Time //到期时间
}

//access_token 对应 appid
var accessTokenCacheList map[string]*accessTokenCache

//access_token 有效时间
var AccessTokenExpires int64

func init() {
	accessTokenCacheList = make(map[string]*accessTokenCache)
	cfg, err := goconfig.LoadConfigFile("etc/creeper.ini")
	if err != nil {
		logrus.Error("goconfig 「etc/creeper.ini」 is error:", err)
		return
	}
	expires, err := cfg.Int64("access_token", "expires")
	if err != nil {
		logrus.Error("cfg.GetValue 「expires」 is error:", err)
		return
	}
	AccessTokenExpires = expires
	//定时清理
	go func() {
		for true {
			cleanRunner()
			//阻塞10秒
			time.Sleep(10 * time.Second)
		}
	}()
}

//设置access_token
func SetAccessToken(accessToken string, appId uint, expires time.Time) error {
	if nil != accessTokenCacheList[accessToken] {
		return errors.New("access_token 不可被重复设置")
	}
	_ = unsetAccessTokenByAppID(appId)
	accessTokenCacheList[accessToken] = &accessTokenCache{AppID: appId, Expires: expires}
	return nil
}

//获取access_token
func GetAccessToken(accessToken string) (appId uint, expires time.Time, err error) {
	if nil == accessTokenCacheList[accessToken] {
		return appId, expires, errors.New("access_token 不存在")
	}
	//验证是否过期
	err = CheckAccessToken(accessToken)
	if err != nil {
		return appId, expires, err
	}
	appId = accessTokenCacheList[accessToken].AppID
	expires = accessTokenCacheList[accessToken].Expires
	return appId, expires, nil
}

//验证过期
func CheckAccessToken(accessToken string) error {
	if nil == accessTokenCacheList[accessToken] {
		return errors.New("无效 access_token")
	}
	//当前时间大于过期时间，清空
	if time.Now().UnixNano() > accessTokenCacheList[accessToken].Expires.UnixNano() {
		delete(accessTokenCacheList, accessToken)
		return errors.New("access_token 已经过期")
	}
	return nil
}

//清理过期无效access_token减少内存开销
func cleanRunner() {
	logrus.Println("开始清理access_token")
	for at := range accessTokenCacheList {
		_ = CheckAccessToken(at)
	}
}

//通过appId清理缓存
func unsetAccessTokenByAppID(appID uint) error {
	for at, ats := range accessTokenCacheList {
		if ats.AppID == appID {
			delete(accessTokenCacheList, at)
			return nil
		}
	}
	logrus.Info("未找到对应AccessToken")
	return errors.New("未找到对应AccessToken")
}
