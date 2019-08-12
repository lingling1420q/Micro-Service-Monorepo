package login

//GetTokenRes get tokrn response
type GetTokenRes struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetSuiteTokenRes get suite token resp
type GetSuiteTokenRes struct {
	// Errcode          int    `json:errcode`
	// Errmsg           string `json:errmsg`
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}
