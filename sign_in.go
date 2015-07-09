package calcutta

import (
	"github.com/adg/xsrftoken"
	"github.com/likestripes/kolkata"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/user/sign_in", SignInFormHandler)
	http.HandleFunc("/user/auth", SignInHandler)
	http.HandleFunc("/user/sign_out", SignOutHandler) //TODO: shouldn't be a GET
}

func SignInFormHandler(w http.ResponseWriter, r *http.Request) {
	scope := kolkata.CreateScope(w, r)
	person, _ := scope.Session()
	if person.Anon {
		param := r.FormValue("param")
		err := Template(w, person, "sign_in", "Please sign in.", param)
		if err != nil {
			scope.Context.Errorf("Template error: " + err.Error())
		}
		return
	}
	http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/error", 307)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	scope := kolkata.CreateScope(w, r)
	person, _ := scope.Session()
	xsrf_token := r.FormValue("xsrf_token")
	xsrf_valid := xsrftoken.Valid(xsrf_token, person.Secret, person.PersonIdStr, "sign_in")

	if person.Anon && xsrf_valid {
		signin := kolkata.SignIn{
			Token:   strings.ToLower(r.FormValue("token")),
			Unsafe:  r.FormValue("password"),
			Context: scope.Context,
		}

		if err := signin.Get(); err == nil {
			if person_id, err := signin.Auth(); person_id != 0 {
				scope.NewSession(person_id)
				http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/hello", 307)
				return
			}else{
				if err != nil {
					scope.Context.Errorf("Auth error: " + err.Error())
				}
			}
		}
	}
	http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/error", 307)
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	scope := kolkata.CreateScope(w, r)
	scope.ClearSession(w, r)
	http.Redirect(scope.Writer, scope.Request, scope.Host, 301)
}
