package main

import (
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
)

// const (
// 	Active = iota + 10
// 	Inactive
// 	Running
// 	Paused
// 	Resumed
// )

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Lshortfile)
	port := os.Getenv("PORT")
	//isFilebased := false
	if port == "" {
		flag.StringVar(&port, "port", "8086", "--port=8086 or -port=8086 or --port 8086 or -port 8086")
		//flag.BoolVar(&isFilebased, "file", false, "--file=true | -file=false")
		flag.Parse()
	}

	log.Println(os.Args)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ICICI Direct"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)

	http.HandleFunc("/user", handlers.CreateUser)

	log.Println("Server started and listening on port ->" + port)
	println("Server started and listening on port ->" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
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
