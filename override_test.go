package methodman

import (
	"gopkg.in/nowk/assert.v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var noop = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	//
})

func TestOverride(t *testing.T) {
	for _, v := range []string{
		"PUT",
		"PATCH",
		"DELETE",
		"POST",
	} {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.PostForm = url.Values{
			"_method": {v},
		}
		w := httptest.NewRecorder()
		MethodOverride(noop).ServeHTTP(w, req)
		assert.Equal(t, v, req.Method)
	}
}

func TestDoesntOverride(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.PostForm = url.Values{
		"_method": {"GET"},
	}
	w := httptest.NewRecorder()
	MethodOverride(noop).ServeHTTP(w, req)
	assert.Equal(t, "POST", req.Method)
}
