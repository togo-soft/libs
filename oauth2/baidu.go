package oauth2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type baidu struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
}

func (this *baidu) GetAccessToken() ([]byte, error) {
	//获取 access_token
	var tokenTarget = fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s", this.code, this.client_id, this.client_secret, this.redirect_uri)
	tokenResp, err := http.Post(tokenTarget, "application/json; charset=utf-8", nil)
	//请求结果
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()
	//结果返回
	return ioutil.ReadAll(tokenResp.Body)
}

func (this *baidu) GetUserInfo(accessToken string) ([]byte, error) {
	//将AccessToken 请求用户信息
	var userInfoTarget = fmt.Sprintf("https://openapi.baidu.com/rest/2.0/passport/users/getInfo?access_token=%s", accessToken)
	_u, err := http.Get(userInfoTarget)
	if err != nil {
		return nil, err
	}
	defer _u.Body.Close()
	return ioutil.ReadAll(_u.Body)
}
