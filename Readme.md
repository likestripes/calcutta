## Calcutta

Readymade sign up/in/out UX for golang.

### Warning

*This is v.01 -- it has no tests, performance is probably awful and it's not proven safe in production anywhere.*  But maybe it'll scratch an itch?

### WTF

`Calcutta` is an opinionated user auth flow built on top of [likestripes/kolkata](https://www.github.com/likestripes/kolkata)

### Install / Import

`go get -u github.com/likestripes/calcutta`

```go
import (
	_ "github.com/likestripes/calcutta"
)
```

### Dependency on `Pacific`

`Kolkata`, and by extension `Calcutta`, uses [likestripes/pacific](https://www.github.com/likestripes/pacific) as an opinionated ORM.  `Pacific` currently supports AppEngine and Postgres.

Google AppEngine: `goapp serve` works out of the box (they include the buildtag for you)

Postgres: `go run -tags 'postgres' main.go` -- details in the [pacific/Readme](https://github.com/likestripes/pacific/blob/master/readme.md).

### Overview

`Calcutta` instantiates simple sign in / sign up forms at:

#### /user/sign_up
![Sign Up](https://baz.likestripes.com/sign_up.png "Sign Up")

#### /user/sign_in
![Sign In](https://baz.likestripes.com/sign_in.png "Sign In")


As well as a sign out endpoint here: `/user/sign_out` *Known issue -- as this is a GET request, a malicious third party could sign out the user via a cross site / iframe src-ing that URL*

Once a user authenticates, they're passed back to either `/user/hello` or `/user/error` via your app.

### `Kolkata` Primer

`Calcutta` establishes a basic sign in flow; to do something with that authenticated user, you'll probably want use `Kolkata`.  `Kolkata` is designed to export a `Person` struct that can be mixed into whatever `WildAndCrazyUser` model your app requires:

```go
type Person struct {
	PersonId    int64
	PersonIdStr string
	Timestamp   time.Time
	Secret      string
	Anon        bool   `datastore:"-" sql:"-" json:"-"`
	Scope       *Scope `datastore:"-" sql:"-" json:"-"`
}
```
The `Person` has _n_ `SignIn`s (a token + password used for authentication) and attaches itself to the current session from your app via `person, err := kolkata.Current(w, r)`.

### Forking, recompiling

The html found in `/templates` gets compiled into `templates.go` via [github.com/mjibson/esc](https://github.com/mjibson/esc)

#### TODO
- [ ] forgot password flow
- [ ] logging
- [ ] documentation!
- [ ] tests!
- [ ] benchmarking

- [x] Contributors welcome!
