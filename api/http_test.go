package api

import (
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"training41/service"
	"training41/store"
)

func TestHTTPCreateAndRead(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	h := New(service.New(d))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("POST", "/", strings.NewReader(`{"ID":"r","TrainingID":"t","UserID":"u"}`)))
	if w.Code != 201 {
		t.Fatalf("code %d", w.Code)
	}
}
