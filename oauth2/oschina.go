package oauth2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type oschina struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
}

func (this *oschina) GetAccessToken() ([]byte, error) {
	var client = &http.Client{}
	var tokenTarget = fmt.Sprintf("https://www.oschina.net/action/openapi/token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s", this.client_id, this.client_secret, this.redirect_uri, this.code)
	tokenReq, _ := http.NewRequest("GET", tokenTarget, nil)
	tokenReq.Header.Add("Content-Type", "application/json")
	tokenReq.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	tokenResp, err := client.Do(tokenReq)
	if err != nil {
		return nil, err
	}
	defer tokenResp.Body.Close()
	return ioutil.ReadAll(tokenResp.Body)
}

func (this *oschina) GetUserInfo(accessToken string) ([]byte, error) {
	//将AccessToken 请求用户信息
	var client = &http.Client{}
	var userInfoTarget = fmt.Sprintf("https://www.oschina.net/action/openapi/user?access_token=%s&dataType=json", accessToken)
	userReq, _ := http.NewRequest("GET", userInfoTarget, nil)
	userReq.Header.Add("Content-Type", "application/json")
	userReq.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	userResq, err := client.Do(userReq)
	if err != nil {
		return nil, err
	}
	defer userResq.Body.Close()
	return ioutil.ReadAll(userResq.Body)
}
