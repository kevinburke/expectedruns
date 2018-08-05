package expectedruns

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Errorf("bad Code: got %d want 200", w.Code)
	}
	if !strings.Contains(w.Body.String(), "65.1%") {
		t.Errorf("expected 65.1%% in the body, got %q", w.Body)
	}
}
