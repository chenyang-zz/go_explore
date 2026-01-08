package concurrence

import (
	"sync"
	"time"
)

var qps = make(chan struct{}, 100)

func handler() {
	qps <- struct{}{}
	defer func() {
		<-qps
	}()
	time.Sleep(3 * time.Second)
}

func QpsLimit() {
	const P = 1000
	var wg sync.WaitGroup
	for i := 0; i < P; i++ {
		wg.Go(handler)
	}
	wg.Wait()
}
