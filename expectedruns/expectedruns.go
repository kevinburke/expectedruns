package expectedruns

import (
	"appengine"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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
	c := appengine.NewContext(r)

	var o Outs
	b, err := ioutil.ReadFile("var/stats.json")
	if err != nil {
		fmt.Fprint(w, "there was an error reading the file: "+err.Error())
		return
	}

	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Fprint(w, "there was an error: "+err.Error())
		return
	}
	b, err = json.Marshal(o)
	c.Infof(string(b))

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
