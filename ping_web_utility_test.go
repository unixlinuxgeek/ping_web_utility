package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInstalled(t *testing.T) {
	err, _ := isInstalled()
	if err != nil {
		t.Fatalf("%s isInstalled() return error!!!", err)
	}
}

func TestMainHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	mainHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected html got: %s\n", err)
	}
}
