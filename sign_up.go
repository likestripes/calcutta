package calcutta

import (
	"github.com/adg/xsrftoken"
	"net/http"
	"github.com/likestripes/kolkata"
)

func init() {
	http.HandleFunc("/user/sign_up", SignUpFormHandler)
	http.HandleFunc("/user/create", SignUpHandler)
}


func SignUpFormHandler(w http.ResponseWriter, r *http.Request) {
	scope := kolkata.CreateScope(w, r)
	person, _ := scope.Session()
	if person.Anon {
		param := r.FormValue("param")
		err := Template(w, person, "sign_up", "Please sign up.", param)
		if err != nil {
			scope.Context.Errorf("Template error: "+err.Error())
		}
		return
	}
	http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/error", 307)
}


func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	scope := kolkata.CreateScope(w, r)
	person, _ := scope.Session()
	xsrf_token := r.FormValue("xsrf_token")
  xsrf_valid := xsrftoken.Valid(xsrf_token, person.Secret, person.PersonIdStr, "sign_up")
  if person.Anon && xsrf_valid {
		password := scope.Request.FormValue("password")

		person := kolkata.Person{
			PersonId: person.PersonId,
      Scope:    &scope,
		}

    sign_ins := map[string]string{
      r.FormValue("username"): password,
      r.FormValue("email"): password,
    }

		person.Save(sign_ins)
		scope.NewSession(person.PersonId)

		http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/hello", 307)
		return
	}

	if !xsrf_valid {
		scope.Context.Errorf("XSRF invalid.")
	}

	http.Redirect(scope.Writer, scope.Request, scope.Host+"/user/error", 307)
}
