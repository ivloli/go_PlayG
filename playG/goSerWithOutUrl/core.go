package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type A struct {
	name string
}

func (s *A) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("Time is: " + tm))
}

func main() {
	a := A{
		name: "A",
	}
	sys_chan := make(chan os.Signal, 1)
	signal.Notify(sys_chan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		http.ListenAndServe("0.0.0.0:23333", &a)
	}()
	<-sys_chan
}
