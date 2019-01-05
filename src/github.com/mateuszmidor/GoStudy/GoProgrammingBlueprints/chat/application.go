package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatarAvatar,
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  Off(),
	}
}

// For google oauth2:
// https://console.developers.google.com/apis/credentials?project=api-project-44918022082 :
// Twój identyfikator klienta
// 		v44918022082-b07tui5r5ud2snbe8ur7ag41qta643ng.apps.googleusercontent.com
// Twój tajny klucz klienta
// 		H3vEKMxxJu-tQjwcE4axM62q

func initOAuth2() {
	gomniauth.SetSecurityKey("AUTH_KEY")
	callbackPrefix := GetAPIEndpoint() + "/auth/callback/"
	gomniauth.WithProviders(google.New("44918022082-b07tui5r5ud2snbe8ur7ag41qta643ng.apps.googleusercontent.com", "H3vEKMxxJu-tQjwcE4axM62q", callbackPrefix+"google"),
		facebook.New("id", "pass", callbackPrefix+"facebook"),
		github.New("id", "pass", callbackPrefix+"github"),
	)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header().Set("Location", "/chat")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// need:
// go get github.com/gorilla/websocket
// go get github.com/stretchr/gomniauth
// go get github.com/stretchr/testify/mock
// go get github.com/clbanning/x2j
// go get github.com/ugorji/go/codec
// go get gopkg.in/mgo.v2/bson
// setup google OAuth client ID under https://console.developers.google.com/apis/credentials?project=api-project-44918022082
// and put the id and pass in "initOAuth2" function
func main() {
	var addr = flag.String("addr", ":"+GetPort(), "Server http address")
	flag.Parse()
	initOAuth2()
	r := newRoom()
	r.tracer = New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", logoutHandler)
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))
	go r.run()
	log.Println("Running the server at", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
