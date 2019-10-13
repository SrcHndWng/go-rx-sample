package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	onNext := handlers.NextFunc(func(item interface{}) {
		fmt.Printf("handled: %v\n", item)
	})

	onError := handlers.ErrFunc(func(err error) {
		fmt.Printf("error: %v\n", err)
	})

	onDone := handlers.DoneFunc(func() {
		fmt.Println("Done!")
	})

	watcher := observer.New(onNext, onError, onDone)

	message := make(chan interface{})
	source := observable.Observable(message)
	sub := source.Subscribe(watcher)

	go func() {
		cnt := 0
		for {
			if cnt > 4 {
				break
			}
			// error sample
			// if cnt == 3 {
			// 	message <- errors.New("some error")
			// } else {
			// 	message <- fmt.Sprintf("Hello, cnt = %d", cnt)
			// }
			message <- fmt.Sprintf("Hello, cnt = %d", cnt) // not error raise. print 'Done!'
			cnt++
			time.Sleep(1 * time.Second)
		}
		close(message)
	}()

	<-sub
}
