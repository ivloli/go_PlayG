package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	jobChan := make(chan Job, 100)
	quit := make(chan struct{})
	var qc int32
	for i := 0; i < 5; i++ {
		go Producer(jobChan, quit)
	}
	for i := 0; i < 10; i++ {
		go Worker(ctx, jobChan, &qc)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(quit)
	cancel()
}

type Job struct {
	Input   int
	Process func(int) int
}

func Producer(out chan<- Job, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
			//product jobs
			j := Job{
				Input: 1,
				Process: func(in int) int {
					fmt.Println(in)
					return in
				},
			}
			out <- j
		}
	}
}

func Worker(ctx context.Context, jobChan <-chan Job, qc *int32) {
	var c int32
	for {
		select {
		case <-ctx.Done():
			return

		case job := <-jobChan:
			atomic.AddInt32(qc, 1)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if c <= atomic.LoadInt32(qc)/10 {
						process(job)
						c++
						break
					}
				}
			}
		}
	}
}

func process(job Job) {
	//do sth with job
}
