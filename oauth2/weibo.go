package oauth2

type weibo struct {
	client_id     string
	client_secret string
	code          string
	state         string
	scope         string
	redirect_uri  string
}
