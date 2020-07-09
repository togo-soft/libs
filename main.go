package main

import "libs/oauth2"

func main() {
	var o = oauth2.Oauth2{
		Source:       "",
		Code:         "",
		ClientID:     "",
		ClientSecret: "",
		RedirectURI:  "",
		Scope:        "",
		State:        "",
	}

	o.NewOauth2()
}
