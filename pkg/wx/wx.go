package wx

import (
	"encoding/json"
	"fmt"
	"github.com/dierbei/blind-box/global"
	"io/ioutil"
	"net/http"
)

//WxSession 微信登陆接口返回session
type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
}

//WxLogin 微信用户授权
func WxLogin(jscode string) (session WxSession, err error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code="+jscode+"&grant_type=%s",
		global.ServerConfig.Wx.AppID,
		global.ServerConfig.Wx.Secret,
		global.ServerConfig.Wx.GrantType,
	)

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return WxSession{}, err
	}

	//处理返回结果
	response, _ := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	jsonStr := string(body)

	//解析json
	if err := json.Unmarshal(body, &session); err != nil {
		session.SessionKey = jsonStr
		return session, err
	}

	return session, err
}
