package main

import (
	"fmt"
	"context"
	"github.com/reactivex/rxgo/v2"
)

type event struct {
	id        int
	eventType string
}

func main() {
	ch := make(chan rxgo.Item, 3)

	go func() {
		ch <- rxgo.Of(event{0, "3D"})
		ch <- rxgo.Of(event{1, "2D"})
		ch <- rxgo.Of(event{2, "2D"})
	}()

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	subscriber2D := observable.Filter(func(i interface{}) bool {
		event := i.(event)
		return event.eventType == "2D"
	})
	subscriber2D.DoOnNext(func(i interface{}) {
		fmt.Printf("STRING 2D: %d\n", i.(event).id)
	})

	subscriber3D := observable.Filter(func(i interface{}) bool {
		event := i.(event)
		return event.eventType == "3D"
	})
	subscriber3D.DoOnNext(func(i interface{}) {
		fmt.Printf("3D: %d\n", i.(event).id)
	})

	observable.Connect(context.Background())
	select {}
}
