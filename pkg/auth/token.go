package auth

import (
    "context"
    "strings"

    "example.com/svc/pkg/util"
)

var log = util.New("auth")

// Validate checks a token. Returns true if valid.
// Token format: "<userid>:<exp-unix>:<sig>".
func Validate(ctx context.Context, token string) bool {
    log.InfoCtx(ctx, "validating token")
    if token == "" {
        return false
    }
    parts := strings.Split(token, ":")
    if len(parts) != 3 {
        return false
    }
    // TODO: verify signature
    return true
}
