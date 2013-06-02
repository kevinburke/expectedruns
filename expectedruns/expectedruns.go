package expectedruns

import (
	//"appengine"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func init() {
	http.HandleFunc("/", root)
}

type Stat struct {
	Id            int // Binary representation of the runners on base
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

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
		return
	}
}

func addPercent(chance float64) string {
	return strconv.FormatFloat(chance, 'g', 3, 64) + "%"
}

func roundDecimal(runs float64) string {
	return strconv.FormatFloat(runs, 'f', 2, 64)
}

func onFirstBase(id int) bool {
	return id&1 == 1
}

func onSecondBase(id int) bool {
	return id&2 == 2
}

func onThirdBase(id int) bool {
	return id&4 == 4
}

func startNewRow(id int) bool {
	return id%2 == 0
}

func root(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	var o Outs
	b, err := ioutil.ReadFile("var/stats.json")
	checkError(err)

	err = json.Unmarshal(b, &o)
	checkError(err)

	b, err = json.Marshal(o)
	checkError(err)
	//c.Infof(string(b))

	t := template.New("master")
	t.Funcs(template.FuncMap{
		"addPercent":   addPercent,
		"roundDecimal": roundDecimal,
		"onFirstBase":  onFirstBase,
		"onSecondBase": onSecondBase,
		"onThirdBase":  onThirdBase,
		"startNewRow":  startNewRow,
	})
	tpl, err := t.ParseFiles("templates/index.html")
	checkError(err)

	err = tpl.Execute(w, o)
	checkError(err)
}
