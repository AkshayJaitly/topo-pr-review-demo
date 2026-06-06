package main

import (
    "net/http"
    "example.com/svc/pkg/auth"
    "example.com/svc/pkg/util"
)

var log = util.New("server")

func main() {
    http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
        tok := r.Header.Get("Authorization")
        if !auth.Validate(tok) {
            http.Error(w, "unauthorized", 401)
            return
        }
        log.Info("ok")
        w.Write([]byte("ok"))
    })
    http.ListenAndServe(":8080", nil)
}
