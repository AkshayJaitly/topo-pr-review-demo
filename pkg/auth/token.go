package auth

import "example.com/svc/pkg/util"

var log = util.New("auth")

// Validate checks a token. Returns true if valid.
func Validate(token string) bool {
    log.Info("validating token")
    return token != ""
}
