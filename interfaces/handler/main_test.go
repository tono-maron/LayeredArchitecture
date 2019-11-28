package handler

import (
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	//レスポンスを生成
	w := httptest.NewRecorder()
}
