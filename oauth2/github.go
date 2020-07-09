package oauth2

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type github struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
}

func (this *github) GetAccessToken() ([]byte, error) {
	//获取 access_token
	var payload = fmt.Sprintf("client_id=%s&client_secret=%s&code=%s&redirect_uri=%s&state=%s", this.client_id, this.client_secret, this.code, this.redirect_uri, this.state)
	tokenResp, err := http.Post("https://github.com/login/oauth/access_token", "application/x-www-form-urlencoded", strings.NewReader(payload))
	//请求结果
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()
	//结果返回
	return ioutil.ReadAll(tokenResp.Body)
}

func (this *github) GetUserInfo(accessToken string) ([]byte, error) {
	//将AccessToken 请求用户信息
	var userInfoTarget = fmt.Sprintf("https://api.github.com/user?access_token=%s", accessToken)
	_u, err := http.Get(userInfoTarget)
	if err != nil {
		return nil, err
	}
	defer _u.Body.Close()
	return ioutil.ReadAll(_u.Body)
}
