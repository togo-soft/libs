package oauth2

// Authorize 认证接口
type Authorize interface {
	GetAccessToken() ([]byte, error)
	GetUserInfo(accessToken string) ([]byte, error)
}

// Oauth2
type Oauth2 struct {
	Source       string `json:"source"`        //类型
	Code         string `json:"code"`          //Code
	ClientID     string `json:"client_id"`     //client_id
	ClientSecret string `json:"client_secret"` //client_secret
	RedirectURI  string `json:"redirect_uri"`  //redirect_uri
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

// NewOauth2 获得对应对象实例
func (this *Oauth2) NewOauth2() Authorize {
	switch this.Source {
	case "github":
		return &github{
			client_id:     this.ClientID,
			client_secret: this.ClientSecret,
			code:          this.Code,
			state:         this.State,
			scope:         this.Scope,
			redirect_uri:  this.RedirectURI,
		}
	case "gitee":
		return &gitee{
			client_id:     this.ClientID,
			client_secret: this.ClientSecret,
			code:          this.Code,
			state:         this.State,
			scope:         this.Scope,
			redirect_uri:  this.RedirectURI,
		}
	case "qq":
		return &qq{
			client_id:     this.ClientID,
			client_secret: this.ClientSecret,
			code:          this.Code,
			state:         this.State,
			scope:         this.Scope,
			redirect_uri:  this.RedirectURI,
		}
	case "oschina":
		return &oschina{
			client_id:     this.ClientID,
			client_secret: this.ClientSecret,
			code:          this.Code,
			state:         this.State,
			scope:         this.Scope,
			redirect_uri:  this.RedirectURI,
		}
	case "baidu":
		return &baidu{
			client_id:     this.ClientID,
			client_secret: this.ClientSecret,
			code:          this.Code,
			state:         this.State,
			scope:         this.Scope,
			redirect_uri:  this.RedirectURI,
		}
	default:
		return nil
	}
}
