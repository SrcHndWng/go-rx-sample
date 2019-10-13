package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
)

func main() {
	f1 := func() interface{} {

		// Simulate a blocking I/O
		time.Sleep(2 * time.Second)
		return 1
	}

	f2 := func() interface{} {

		// Simulate a blocking I/O
		time.Sleep(time.Second)
		return 2
	}

	f3 := func() interface{} {

		// Simulate a blocking I/O
		time.Sleep(3 * time.Second)
		return 3
	}

	f4 := func() interface{} {

		// Simulate a blocking I/O
		time.Sleep(4 * time.Second)
		return errors.New("some error")
	}

	onNext := handlers.NextFunc(func(v interface{}) {
		fmt.Printf("handled, v = %v\n", v)
	})

	wait := observable.Start(f1, f2, f3, f4).Subscribe(onNext)
	sub := <-wait

	if err := sub.Err(); err != nil {
		fmt.Printf("error raise. err = %v\n", err)
	}

	fmt.Println("finish")
}
