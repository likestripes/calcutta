package calcutta

import (
	"bytes"
	"github.com/adg/xsrftoken"
	"github.com/likestripes/kolkata"
	"html/template"
	"io/ioutil"
	"net/http"
)

type TemplateData map[string]interface{}

var fs = FS(false)

func Template(w http.ResponseWriter, person kolkata.Person, tpl_name, page_title, param string) error {

	var inner bytes.Buffer
	xsrf_token := xsrftoken.Generate(person.Secret, person.PersonIdStr, tpl_name)
	fields := TemplateData{
		"xsrf_token": xsrf_token,
		"param":      param,
	}

	tpl_file, err := file("/" + tpl_name + "_form.html")

	if err != nil {
		return err
	}

	tpl, err := template.New(tpl_name).Parse(tpl_file)

	err = tpl.Execute(&inner, fields)

	if err != nil {
		return err
	}

	form := TemplateData{
		"form":       template.HTML(inner.String()),
		"page_title": page_title,
	}

	base_file, err := file("/base.html")

	if err != nil {
		return err
	}

	base, err := template.New("base").Parse(base_file)

	if err != nil {
		return err
	}

	return base.Execute(w, form)
}

func file(file_name string) (contents string, err error) {
	f, err := fs.Open(file_name)
	if err != nil {
		return "", err
	}
	t, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	f.Close()
	return string(t), nil
}

//go:generate esc -o templates.go -pkg calcutta -prefix templates templates/base.html templates/sign_in_form.html templates/sign_up_form.html
