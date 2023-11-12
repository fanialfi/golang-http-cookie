package main

import (
	"log"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]any

var (
	cookieName = "CookieData"
	port       = ":8080"
)

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)
	http.HandleFunc("/ok", ActionOk)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{
			Name:    cookieName,
			Value:   gubrak.RandomString(32),
			Expires: time.Now().Add(time.Minute),
		}
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:    cookieName,
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func ActionOk(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Okay"))
}
