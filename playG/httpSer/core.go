package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	fmt.Println("vim-go")
	limiter := rate.NewLimiter(rate.Every(time.Second), 5)
	mux := http.NewServeMux()
	allowHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("%+v", r)
		err := waitToken(limiter)
		if err != nil {
			w.Write([]byte(`{"code":"-1","msg":"` + err.Error() + `"}`))
		} else {
			w.Write([]byte(`{"code":0,"msg":"ok"}`))
		}
	}
	mux.HandleFunc("/get_allow", allowHandlerFunc)
	go func() {
		http.ListenAndServe(":18080", mux)
	}()
	go func() {
		for {
			resp, err := http.DefaultClient.Get("http://127.0.0.1:18080/get_allow")
			b, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(b))
			fmt.Println(err)
			time.Sleep(time.Second)
		}
	}()
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
}

func waitToken(l *rate.Limiter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	e := l.Wait(ctx)
	return e
}
