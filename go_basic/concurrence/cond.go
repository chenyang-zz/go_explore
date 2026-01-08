package concurrence

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func CondSignal() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	var wg sync.WaitGroup

	wg.Go(func() {
		time.Sleep(time.Second)
		cond.Signal() //只有一个协程能收到信号
		fmt.Println("发出信号")
		time.Sleep(time.Second)
		cond.Signal()
		fmt.Println("发出信号")
	})

	wg.Go(func() {
		cond.L.Lock() //Wait()函数内部会先执行Unlock()，所以在调Wait()之前需要先调Lock()
		cond.Wait()   //阻塞，等信号。
		fmt.Println("收到信号，执行某些工作")
		cond.L.Unlock() //释放锁，因为下次Wait()前还得拿到锁

		// time.Sleep(2 * time.Second) // 加上这行代码，则Wait()永远等不到信号
		cond.L.Lock()
		cond.Wait() // 必须先调Wait()，之后再调Signal()或Broadcast()，这样Wait()才能解除阻塞
		fmt.Println("收到信号，执行某些工作")
		cond.L.Unlock()
	})

	wg.Wait()
	fmt.Println(strings.Repeat("-", 50))
}

func ChannelSignal() {
	ch := make(chan struct{}, 100)

	var wg sync.WaitGroup
	wg.Go(func() {
		time.Sleep(time.Second)
		ch <- struct{}{}
		fmt.Println("发出信号")
		time.Sleep(time.Second)
		ch <- struct{}{}
		fmt.Println("发出信号")
	})

	wg.Go(func() {
		<-ch
		fmt.Println("收到信号，执行某些工作")
		time.Sleep(2 * time.Second)
		<-ch
		fmt.Println("收到信号，执行某些工作")
	})

	wg.Wait()
	fmt.Println(strings.Repeat("-", 50))
}

func CondBroadcast() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	var wg sync.WaitGroup
	wg.Go(func() {
		time.Sleep(time.Second)
		cond.Broadcast() //所有协程都能收到信号
		fmt.Println("广播信号")

		time.Sleep(time.Second)
		cond.Broadcast()
		fmt.Println("广播信号")
	})

	for i := 0; i < 3; i++ {
		wg.Go(func() {
			cond.L.Lock()
			cond.Wait()
			fmt.Println("收到信号，执行某些工作")
			cond.L.Unlock()

			// time.Sleep(2 * time.Second)
			cond.L.Lock()
			cond.Wait()
			fmt.Println("收到信号，执行某些工作")
			cond.L.Unlock()
		})

	}
	wg.Wait()
	fmt.Println(strings.Repeat("-", 50))
}

func ChannelBroadcast() {
	ch := make(chan struct{})

	var wg sync.WaitGroup

	wg.Go(func() {
		time.Sleep(time.Second)
		close(ch) // channel只能close一次，所以只能广播一次；如果想反复广播需要使用sync.Cond
		fmt.Println("广播信号")
	})

	for i := 0; i < 3; i++ {
		wg.Go(func() {
			time.Sleep(2 * time.Second)
			<-ch
			fmt.Println("收到信号，执行某些工作")
		})
	}
	wg.Wait()
	fmt.Println(strings.Repeat("-", 50))
}
