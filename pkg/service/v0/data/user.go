package data

// User holds the payload for a GetUser response
type User struct {
	// TODO needs better naming, clarify if we need a userid, a username or both
	UserID      string `json:"id" xml:"id"`
	Username    string `json:"username" xml:"username"`
	DisplayName string `json:"displayname" xml:"displayname"`
	Email       string `json:"email" xml:"email"`
	Enabled     bool   `json:"enabled" xml:"enabled"`
}

// Users holds user ids for the user listing
type Users struct {
	Users []string `json:"users" xml:"users>element"`
}

// SigningKey holds the Payload for a GetSigningKey response
type SigningKey struct {
	User       string `json:"user" xml:"user"`
	SigningKey string `json:"signing-key" xml:"signing-key"`
}
