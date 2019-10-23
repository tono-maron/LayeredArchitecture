package main

import (
	"net/http"
	"flag"
	"log"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	/* ===== URLマッピングを行う ===== */
	// http.Handle("/user/get",
	// 	get(middleware.Authenticate(handler.HandleUserGet())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}