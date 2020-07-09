package oauth2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type qq struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
	open_id       string
}

func (this *qq) GetAccessToken() ([]byte, error) {
	//获取 access_token
	var tokenTarget = fmt.Sprintf("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&fmt=json&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s", this.client_id, this.client_secret, this.code, this.redirect_uri)
	tokenResp, err := http.Get(tokenTarget)
	//请求结果
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()
	//结果返回
	return ioutil.ReadAll(tokenResp.Body)
}

func (this *qq) GetUserInfo(accessToken string) ([]byte, error) {
	//将AccessToken 请求用户信息
	var userInfoTarget = fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", accessToken, this.client_id, this.open_id)
	_u, err := http.Get(userInfoTarget)
	if err != nil {
		return nil, err
	}
	defer _u.Body.Close()
	return ioutil.ReadAll(_u.Body)
}
