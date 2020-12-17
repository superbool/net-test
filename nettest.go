package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var mode = flag.String("mode", "client", "mode为server或者client")
var host = flag.String("host", "localhost", "服务器ip")

func main() {
	flag.Parse()
	if *mode == "server" {
		server()
	} else {
		client(*host)
	}
}

func server() {
	log.Println("start server :8000")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\n", r.RemoteAddr)
	w.Write([]byte("hello world"))
}

func client(host string) {
	start := now()
	resp, err := http.Get("http://" + host + ":8000/")
	end := now()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("耗时毫秒:%v,微秒:%v,纳秒:%v\n响应:%s\n", (end-start)/1000000, (end-start)/1000, (end - start), body)
}

func now() int64 {
	return time.Now().UnixNano()
}
