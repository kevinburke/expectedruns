package expectedruns

import (
	//"appengine"

	"bytes"
	"html/template"
	"io"
	"net/http"
	"runtime"
	"strconv"

	expectedruns "github.com/kevinburke/expectedruns/lib"
)

func init() {
	http.HandleFunc("/", root)
}

func addPercent(chance float64) string {
	return strconv.FormatFloat(chance, 'g', 3, 64) + "%"
}

func roundDecimal(runs float64) string {
	return strconv.FormatFloat(runs, 'f', 2, 64)
}

// XXX, investigate doing something neat with partials to turn these into one
// function.
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

var tpl *template.Template
var td *templateData

func init() {
	t := template.New("master")
	t.Funcs(template.FuncMap{
		"addPercent":   addPercent,
		"roundDecimal": roundDecimal,
		"onFirstBase":  onFirstBase,
		"onSecondBase": onSecondBase,
		"onThirdBase":  onThirdBase,
		"startNewRow":  startNewRow,
	})
	var err error
	tpl, err = t.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	td = &templateData{Version: runtime.Version(), Stats: expectedruns.Stats}
}

type templateData struct {
	Version string
	Stats   []expectedruns.Out
}

func root(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, td)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "An error occurred")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, buf)
}
