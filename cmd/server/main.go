package main

import (
    "context"
    "net/http"

    "example.com/svc/pkg/auth"
    "example.com/svc/pkg/util"
)

var log = util.New("server")

func main() {
    http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()
        tok := r.Header.Get("Authorization")
        if !auth.Validate(ctx, tok) {
            http.Error(w, "unauthorized", 401)
            return
        }
        log.InfoCtx(ctx, "ok")
        w.Write([]byte("ok"))
    })
    http.ListenAndServe(":8080", nil)
}
