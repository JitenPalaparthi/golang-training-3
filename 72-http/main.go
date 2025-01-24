package main

import (
	"demo/handlers"
	"log"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ICICI Direct"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)

	http.HandleFunc("/user", handlers.CreateUser)

	log.Println("Server started and listening on port 8086")
	if err := http.ListenAndServe(":8086", nil); err != nil {
		log.Println(err.Error())
		runtime.Goexit()
	} // 0.0.0.0:8086
}

// how many interfaces ? 0.0.0.0 --? 127.0.0.1 -> local network
// 192.168.0.1 ->
// thuderbolt ports, wifi, bluthooth
// switch -->
// router
// anp
