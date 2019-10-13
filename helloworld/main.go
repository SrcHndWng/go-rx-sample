package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Printf("handled: %v\n", item)
		},

		ErrHandler: func(err error) {
			fmt.Printf("error: %v\n", err)
		},

		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

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
