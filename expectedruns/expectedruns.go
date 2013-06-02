package expectedruns

import (
	//"appengine"
	"fmt"
	"html/template"
	"net/http"
	//"path/filepath"
)

func init() {
	http.HandleFunc("/", root)
}

type Stat struct {
	Id            int
	ScoringChance float64
	ExpectedRuns  float64
}

type Out struct {
	FriendlyName string
	Stats        []Stat
}

type Outs struct {
	Outs []Out
}

func root(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	//c.Infof(path)
	t := template.New("master")
	tpl, err := t.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "there was an error: "+err.Error())
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "there was an error: "+err.Error())
		return
	}
}
