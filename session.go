package domain

// Session holds on to the token for accessing API
type Session struct {
	accessToken string
}

// NewSession is the main way to consume this package
func NewSession(clientID string, secret string) (*Session, error) {
	session := new(Session)
	token, err := getAccessToken(clientID, secret)
	if err != nil {
		return nil, err
	}
	session.accessToken = token
	return session, nil
}
