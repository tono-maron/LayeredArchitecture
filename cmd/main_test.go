package main

func TestMain() {
	router := NewRouter()

    req := httptest.NewRequest("GET", "/hello", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)
}
