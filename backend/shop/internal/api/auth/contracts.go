package authapi

type LoginResponse struct {
	AuthorizeURL string `json:"authorize_url"`
}

type CallbackResponse struct {
	SessionToken string   `json:"session_token"`
	User         AppUser  `json:"user"`
	Provider     string   `json:"provider"`
	Scopes       []string `json:"scopes"`
}

type AppUser struct {
	ID          int64  `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Balance     int    `json:"balance"`
}

type OAuthState struct {
	State string
}

type OAuthUser struct {
	ProviderID  int64
	Login       string
	DisplayName string
	Email       string
}

type FortyTwoTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type FortyTwoUserResponse struct {
	ID          int64  `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
}
