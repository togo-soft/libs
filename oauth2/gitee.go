package oauth2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type gitee struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
}

func (this *gitee) GetAccessToken() ([]byte, error) {
	//获取 access_token
	var tokenTarget = fmt.Sprintf("https://gitee.com/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s&scope=%s&state=%s", this.code, this.client_id, this.redirect_uri, this.client_secret, this.scope, this.state)
	tokenResp, err := http.Post(tokenTarget, "application/json; charset=utf-8", nil)
	//请求结果
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()
	//结果返回
	return ioutil.ReadAll(tokenResp.Body)
}

func (this *gitee) GetUserInfo(accessToken string) ([]byte, error) {
	//将AccessToken 请求用户信息
	var userInfoTarget = fmt.Sprintf("https://gitee.com/api/v5/user?access_token=%s", accessToken)
	_u, err := http.Get(userInfoTarget)
	if err != nil {
		return nil, err
	}
	defer _u.Body.Close()
	return ioutil.ReadAll(_u.Body)
}
